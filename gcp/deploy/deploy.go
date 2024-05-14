package deploy

import (
	mongodb "github.com/nitrictech/mongodb-provider/common/deploy"
	common "github.com/nitrictech/nitric/cloud/common/deploy"
	"github.com/nitrictech/nitric/cloud/common/deploy/pulumix"
	"github.com/nitrictech/nitric/cloud/gcp/deploy"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GcpExtensionProvider struct {
	deploy.NitricGcpPulumiProvider
	mongodb.MongoDBProvider
}

func NewGcpExtensionProvider() *GcpExtensionProvider {
	gcpProvider := deploy.NewNitricGcpProvider()

	mongoProvider := mongodb.NewMongoDBProvider("GCP")

	return &GcpExtensionProvider{
		NitricGcpPulumiProvider: *gcpProvider,
		MongoDBProvider:         *mongoProvider,
	}
}

func (a *GcpExtensionProvider) Config() (auto.ConfigMap, error) {
	config, err := a.NitricGcpPulumiProvider.Config()
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

func (a *GcpExtensionProvider) Init(attributes map[string]interface{}) error {
	var err error

	a.CommonStackDetails, err = common.CommonStackDetailsFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	a.GcpConfig, err = deploy.ConfigFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Bad stack configuration: %s", err)
	}

	a.MongoDBConfig, err = mongodb.ConfigFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Bad stack configuration: %s", err)
	}

	return nil
}

func (a *GcpExtensionProvider) Pre(ctx *pulumi.Context, resources []*pulumix.NitricPulumiResource[any]) error {
	err := a.NitricGcpPulumiProvider.Pre(ctx, resources)
	if err != nil {
		return err
	}

	err = a.MongoDBProvider.Pre(ctx, resources, a.ProjectName, a.Region)
	if err != nil {
		return err
	}

	return nil
}
