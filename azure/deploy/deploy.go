package deploy

import (
	mongodb "github.com/nitrictech/mongodb-provider/common/deploy"
	"github.com/nitrictech/nitric/cloud/azure/deploy"
	common "github.com/nitrictech/nitric/cloud/common/deploy"
	"github.com/nitrictech/nitric/cloud/common/deploy/pulumix"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AzureExtensionProvider struct {
	deploy.NitricAzurePulumiProvider
	mongodb.MongoDBProvider
}

func NewAzureExtensionProvider() *AzureExtensionProvider {
	azureProvider := deploy.NewNitricAzurePulumiProvider()

	mongoProvider := mongodb.NewMongoDBProvider("AZURE")

	return &AzureExtensionProvider{
		NitricAzurePulumiProvider: *azureProvider,
		MongoDBProvider:           *mongoProvider,
	}
}

func (a *AzureExtensionProvider) Config() (auto.ConfigMap, error) {
	config, err := a.NitricAzurePulumiProvider.Config()
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

func (a *AzureExtensionProvider) Init(attributes map[string]interface{}) error {
	var err error

	a.CommonStackDetails, err = common.CommonStackDetailsFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	a.AzureConfig, err = deploy.ConfigFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Bad stack configuration: %s", err)
	}

	a.MongoDBConfig, err = mongodb.ConfigFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Bad stack configuration: %s", err)
	}

	return nil
}

func (a *AzureExtensionProvider) Pre(ctx *pulumi.Context, resources []*pulumix.NitricPulumiResource[any]) error {
	err := a.NitricAzurePulumiProvider.Pre(ctx, resources)
	if err != nil {
		return err
	}

	err = a.MongoDBProvider.Pre(ctx, resources, a.ProjectName, a.Region)
	if err != nil {
		return err
	}

	return nil
}
