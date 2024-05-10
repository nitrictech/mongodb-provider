package deploy

import (
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	resourcespb "github.com/nitrictech/nitric/core/pkg/proto/resources/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/samber/lo"
)

func (a *GcpExtensionProvider) Policy(ctx *pulumi.Context, parent pulumi.Resource, name string, config *deploymentspb.Policy) error {
	filteredConfig := deploymentspb.Policy{
		Principals: config.Principals,
	}

	filteredConfig.Resources = lo.Filter(config.Resources, func(res *deploymentspb.Resource, idx int) bool {
		return res.Id.Type != resourcespb.ResourceType_KeyValueStore
	})

	filteredConfig.Actions = lo.Filter(config.Actions, func(res resourcespb.Action, idx int) bool {
		return !lo.Contains([]resourcespb.Action{
			resourcespb.Action_KeyValueStoreDelete,
			resourcespb.Action_KeyValueStoreRead,
			resourcespb.Action_KeyValueStoreWrite,
		}, res)
	})

	if len(filteredConfig.Actions) == 0 {
		return nil
	}

	return a.NitricGcpPulumiProvider.Policy(ctx, parent, name, &filteredConfig)
}
