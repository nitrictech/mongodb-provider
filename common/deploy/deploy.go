package deploy

import (
	"fmt"
	"os"
	"strings"

	"github.com/nitrictech/nitric/cloud/common/deploy/pulumix"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/pulumi/pulumi-mongodbatlas/sdk/v2/go/mongodbatlas"
	mongodb "github.com/pulumi/pulumi-mongodbatlas/sdk/v2/go/mongodbatlas"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/samber/lo"
)

type MongoDBProvider struct {
	MongoDBConfig *MongoDBConfig
	Provider      string
}

func NewMongoDBProvider(provider string) *MongoDBProvider {
	return &MongoDBProvider{
		Provider: provider,
	}
}

func (p *MongoDBProvider) Pre(ctx *pulumi.Context, resources []*pulumix.NitricPulumiResource[any], projectName string, region string) error {
	// Check if a key value store exists, if so get/create a (default) firestore database
	databases := lo.Filter(resources, func(res *pulumix.NitricPulumiResource[any], idx int) bool {
		_, ok := res.Config.(*deploymentspb.Resource_KeyValueStore)
		return ok
	})

	if len(databases) > 0 {
		project, err := mongodb.NewProject(ctx, projectName, &mongodb.ProjectArgs{
			Name:  pulumi.String(projectName),
			OrgId: pulumi.String(p.MongoDBConfig.OrgId),
		})
		if err != nil {
			return err
		}

		clusterName := "nitric"

		cluster, err := mongodbatlas.NewCluster(ctx, clusterName, &mongodbatlas.ClusterArgs{
			ProjectId:                project.ID(),
			Name:                     pulumi.String(clusterName),
			AutoScalingDiskGbEnabled: pulumi.Bool(true),
			ProviderRegionName:       pulumi.String(strings.ToUpper(region)),
			MongoDbMajorVersion:      pulumi.String("7.0"),
			ProviderName:             pulumi.String("TENANT"),
			BackingProviderName:      pulumi.String(p.Provider),
			ProviderInstanceSizeName: pulumi.String("M0"),
		})
		if err != nil {
			return err
		}

		// generate a db cluster random password
		dbMasterPassword, err := random.NewRandomPassword(ctx, "db-master-password", &random.RandomPasswordArgs{
			Length:  pulumi.Int(16),
			Special: pulumi.Bool(false),
		})
		if err != nil {
			return err
		}

		user, err := mongodb.NewDatabaseUser(ctx, "nitric-user", &mongodb.DatabaseUserArgs{
			Username:         pulumi.String("nitric-user"),
			Password:         dbMasterPassword.Result,
			ProjectId:        project.ID(),
			AuthDatabaseName: pulumi.String("admin"),
			Roles: mongodbatlas.DatabaseUserRoleArray{
				&mongodbatlas.DatabaseUserRoleArgs{
					RoleName:     pulumi.String("readWriteAnyDatabase"),
					DatabaseName: pulumi.String("admin"),
				},
			},
		})
		if err != nil {
			return err
		}

		_, err = mongodbatlas.NewProjectIpAccessList(ctx, "nitric-ip", &mongodbatlas.ProjectIpAccessListArgs{
			ProjectId: project.ID(),
			CidrBlock: pulumi.String("0.0.0.0/0"),
			Comment:   pulumi.String("cidr block for lambda, container app, or cloud run access"),
		})
		if err != nil {
			return err
		}

		clusterUrl := cluster.SrvAddress.ApplyT(func(uri string) string {
			// Get the server address without the substring e.g "mongodb+srv://id.mongodb.net"
			return uri[14:]
		}).(pulumi.StringOutput)

		// append the mongodb environment variables to all the services
		for _, res := range resources {
			config, ok := res.Config.(*pulumix.NitricPulumiServiceConfig)

			if ok {
				clusterUrl := pulumi.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=nitric", user.Username, dbMasterPassword.Result, clusterUrl)

				config.SetEnv("MONGO_CLUSTER_CONNECTION_STRING", clusterUrl)
				config.SetEnv("MONGODB_ATLAS_PRIVATE_KEY", nil)
				config.SetEnv("MONGODB_ATLAS_PUBLIC_KEY", nil)
			}
		}
	}

	return nil
}

func (p *MongoDBProvider) MongoConfig() (auto.ConfigMap, error) {
	publicKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	if publicKey == "" {
		return auto.ConfigMap{}, fmt.Errorf("error getting MONGODB_ATLAS_PUBLIC_KEY, has it been set?")
	}

	privateKey := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")
	if privateKey == "" {
		return auto.ConfigMap{}, fmt.Errorf("error getting MONGODB_ATLAS_PRIVATE_KEY, has it been set?")
	}

	mongoVersion := "3.15.0"

	return auto.ConfigMap{
		"mongodbatlas:publicKey":  auto.ConfigValue{Value: publicKey, Secret: true},
		"mongodbatlas:privateKey": auto.ConfigValue{Value: privateKey, Secret: true},
		"mongodbatlas:version":    auto.ConfigValue{Value: mongoVersion},
	}, nil
}
