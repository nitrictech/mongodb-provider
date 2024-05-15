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

var AzureAtlasRegionMap = map[string]string{
	// Americas
	"centralus":       "US_CENTRAL",
	"eastus":          "US_EAST",
	"eastus2":         "US_EAST_2",
	"northcentralus":  "US_NORTH_CENTRAL",
	"southcentralus":  "US_SOUTH_CENTRAL",
	"westus":          "US_WEST",
	"westus2":         "US_WEST_2",
	"canadacentral":   "CANADA_CENTRAL",
	"canadaeast":      "CANADA_EAST",
	"brazilsouth":     "BRAZIL_SOUTH",
	"brazilsoutheast": "BRAZIL_SOUTHEAST",
	"westcentralus":   "US_WEST_CENTRAL",
	"westus3":         "US_WEST_3",

	// Europe
	"northeurope":        "EUROPE_NORTH",
	"westeurope":         "EUROPE_WEST",
	"uksouth":            "UK_SOUTH",
	"ukwest":             "UK_WEST",
	"francecentral":      "FRANCE_CENTRAL",
	"francesouth":        "FRANCE_SOUTH",
	"italynorth":         "ITALY_NORTH",
	"germanywestcentral": "GERMANY_WEST_CENTRAL",
	"germanynorth":       "GERMANY_NORTH",
	"polandcentral":      "POLAND_CENTRAL",
	"switzerlandnorth":   "SWITZERLAND_NORTH",
	"switzerlandwest":    "SWITZERLAND_WEST",
	"norwayeast":         "NORWAY_EAST",
	"norwaywest":         "NORWAY_WEST",
	"sweedencentral":     "SWEEDEN_CENTRAL",
	"swedensouth":        "SWEEDEN_SOUTH",

	// Asia Pacific
	"eastasia":           "ASIA_EAST",
	"southeastasia":      "ASIA_SOUTH_EAST",
	"australiacentral":   "AUSTRALIA_CENTRAL",
	"australiacentral2":  "AUSTRALIA_CENTRAL_2",
	"australiaeast":      "AUSTRALIA_EAST",
	"australiasoutheast": "AUSTRALIA_SOUTH_EAST",
	"centralindia":       "INDIA_CENTRAL",
	"southindia":         "INDIA_SOUTH",
	"westindia":          "INDIA_WEST",
	"japaneast":          "JAPAN_EAST",
	"japanwest":          "JAPAN_WEST",
	"koreacentral":       "KOREA_CENTRAL",
	"koreasouth":         "KOREA_SOUTH",

	// Africa
	"southafricanorth": "SOUTH_AFRICA_NORTH",
	"southafricawest":  "SOUTH_AFRICA_WEST",

	// Middle East
	"uaecentral":    "UAE_CENTRAL",
	"uaenorth":      "UAE_NORTH",
	"qatarcentral":  "QATAR_CENTRAL",
	"israelcentral": "ISRAEL_CENTRAL",
}

func (a *AzureExtensionProvider) Pre(ctx *pulumi.Context, resources []*pulumix.NitricPulumiResource[any]) error {
	err := a.NitricAzurePulumiProvider.Pre(ctx, resources)
	if err != nil {
		return err
	}

	atlasRegion, ok := AzureAtlasRegionMap[a.Region]
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Unsupported mongo atlas region %s", a.Region)
	}

	err = a.MongoDBProvider.Pre(ctx, resources, a.ProjectName, atlasRegion)
	if err != nil {
		return err
	}

	return nil
}
