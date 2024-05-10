package deploy

import (
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	"github.com/nitrictech/nitric/cloud/common/deploy/pulumix"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func (a *AzureExtensionProvider) Service(ctx *pulumi.Context, parent pulumi.Resource, name string, config *pulumix.NitricPulumiServiceConfig, runtime provider.RuntimeProvider) error {
	config.SetEnv("MONGO_CLUSTER_CONNECTION_STRING", a.ClusterURL)

	return a.NitricAzurePulumiProvider.Service(ctx, parent, name, config, runtime)
}
