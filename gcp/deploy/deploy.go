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

var GcpAtlasRegionMap = map[string]string{
	// Americas
	"us-central1":             "CENTRAL_US",
	"us-east1":                "EASTERN_US",
	"us-east4":                "US_EAST_4",
	"us-west5":                "US_EAST_5",
	"northamerica-northeast1": "NORTH_AMERICA_NORTHEAST_1",
	"northamerica-northeast2": "NORTH_AMERICA_NORTHEAST_2",
	"southamerica-east1":      "SOUTH_AMERICA_EAST_1",
	"southamerica-west1":      "SOUTH_AMERICA_WEST_1",
	"us-west1":                "WESTERN_US",
	"us-west2":                "US_WEST_2",
	"us-west3":                "US_WEST_3",
	"us-west4":                "US_WEST_4",
	"us-south1":               "US_SOUTH_1",

	// Asia Pacific
	"asia-east1":           "EASTERN_ASIA_PACIFIC",
	"asia-east2":           "ASIA_EAST_2",
	"asia-northeast1":      "NORTHEASTERN_ASIA_PACIFIC",
	"asia-northeast2":      "ASIA_NORTHEAST_2",
	"asia-northeast3":      "ASIA_NORTHEAST_3",
	"asia-south1":          "SOUTHERN_ASIA_PACIFIC",
	"asia-southeast1":      "SOUTHEASTERN_ASIA_PACIFIC",
	"asia-southeast2":      "ASIA_SOUTHEAST_2",
	"australia-southeast1": "AUSTRALIA_SOUTHEAST_1",
	"australia-southeast2": "AUSTRALIA_SOUTHEAST_2",

	// Europe
	"europe-central2":   "EUROPE_CENTRAL_2",
	"europe-north1":     "EUROPE_NORTH_1",
	"europe-west2":      "EUROPE_WEST_2",
	"europe-west3":      "EUROPE_WEST_3",
	"europe-west4":      "EUROPE_WEST_4",
	"europe-west6":      "EUROPE_WEST_6",
	"europe-west10":     "EUROPE_WEST_10",
	"europe-west1":      "WESTERN_EUROPE",
	"europe-west9":      "EUROPE_WEST_9",
	"europe-west12":     "EUROPE_WEST_12",
	"europe-southwest1": "EUROPE_SOUTHWEST_1",

	// Middle East
	"me-west1":    "MIDDLE_EAST_WEST_1",
	"me-central1": "MIDDLE_EAST_CENTRAL_1",
	"me-central2": "MIDDLE_EAST_CENTRAL_2",
}

func (a *GcpExtensionProvider) Pre(ctx *pulumi.Context, resources []*pulumix.NitricPulumiResource[any]) error {
	err := a.NitricGcpPulumiProvider.Pre(ctx, resources)
	if err != nil {
		return err
	}

	atlasRegion, ok := GcpAtlasRegionMap[a.Region]
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Unsupported mongo atlas region %s", a.Region)
	}

	err = a.MongoDBProvider.Pre(ctx, resources, a.ProjectName, atlasRegion)
	if err != nil {
		return err
	}

	return nil
}
