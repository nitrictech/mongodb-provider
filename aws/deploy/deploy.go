package deploy

import (
	"strings"

	mongodb "github.com/nitrictech/mongodb-provider/common/deploy"
	"github.com/nitrictech/nitric/cloud/aws/deploy"
	common "github.com/nitrictech/nitric/cloud/common/deploy"
	"github.com/nitrictech/nitric/cloud/common/deploy/pulumix"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AwsExtensionProvider struct {
	deploy.NitricAwsPulumiProvider
	mongodb.MongoDBProvider
}

func NewAwsExtensionProvider() *AwsExtensionProvider {
	awsProvider := deploy.NewNitricAwsProvider()

	mongoProvider := mongodb.NewMongoDBProvider("AWS")

	return &AwsExtensionProvider{
		NitricAwsPulumiProvider: *awsProvider,
		MongoDBProvider:         *mongoProvider,
	}
}

func (a *AwsExtensionProvider) Config() (auto.ConfigMap, error) {
	config, err := a.NitricAwsPulumiProvider.Config()
	if err != nil {
		return nil, err
	}

	mongoConfig, err := a.MongoDBProvider.MongoConfig()
	if err != nil {
		return nil, err
	}

	for k, v := range mongoConfig {
		config[k] = v
	}

	return config, nil
}

func (a *AwsExtensionProvider) Init(attributes map[string]interface{}) error {
	var err error

	a.CommonStackDetails, err = common.CommonStackDetailsFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	a.AwsConfig, err = deploy.ConfigFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Bad stack configuration: %s", err)
	}

	a.MongoDBConfig, err = mongodb.ConfigFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Bad stack configuration: %s", err)
	}

	return nil
}

func translateRegion(region string) string {
	return strings.ToUpper(strings.Replace(region, "-", "_", -1))
}

func (a *AwsExtensionProvider) Pre(ctx *pulumi.Context, resources []*pulumix.NitricPulumiResource[any]) error {
	err := a.NitricAwsPulumiProvider.Pre(ctx, resources)
	if err != nil {
		return err
	}

	err = a.MongoDBProvider.Pre(ctx, resources, a.ProjectName, translateRegion(a.Region))
	if err != nil {
		return err
	}

	return nil
}
