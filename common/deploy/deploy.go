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
	Project *mongodbatlas.Project
	Cluster *mongodbatlas.Cluster

	MongoDBConfig *MongoDBConfig
	Provider      string
	ClusterURL    pulumi.StringOutput
}

func NewMongoDBProvider(provider string) *MongoDBProvider {
	return &MongoDBProvider{
		Provider: provider,
	}
}

func (p *MongoDBProvider) Pre(ctx *pulumi.Context, resources []*pulumix.NitricPulumiResource[any], projectName string, region string) error {
	var err error

	// Check if a key value store exists, if so get/create a (default) firestore database
	databases := lo.Filter(resources, func(res *pulumix.NitricPulumiResource[any], idx int) bool {
		_, ok := res.Config.(*deploymentspb.Resource_KeyValueStore)
		return ok
	})

	if len(databases) > 0 {
		p.Project, err = mongodb.NewProject(ctx, projectName, &mongodb.ProjectArgs{
			Name:  pulumi.String(projectName),
			OrgId: pulumi.String(p.MongoDBConfig.OrgId),
		})
		if err != nil {
			return err
		}

		clusterName := "nitric"

		p.Cluster, err = mongodbatlas.NewCluster(ctx, clusterName, &mongodbatlas.ClusterArgs{
			ProjectId:                p.Project.ID(),
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

		roles := mongodbatlas.DatabaseUserRoleArray{
			&mongodbatlas.DatabaseUserRoleArgs{
				RoleName:     pulumi.String("readAnyDatabase"),
				DatabaseName: pulumi.String("admin"),
			},
		}

		for _, db := range databases {
			dbName := db.Id.Name

			roles = append(roles, &mongodbatlas.DatabaseUserRoleArgs{
				RoleName:     pulumi.String("readWrite"),
				DatabaseName: pulumi.String(dbName),
			})
		}

		user, err := mongodb.NewDatabaseUser(ctx, "nitric-user", &mongodb.DatabaseUserArgs{
			Username:         pulumi.String("nitric-user"),
			Password:         dbMasterPassword.Result,
			ProjectId:        p.Project.ID(),
			AuthDatabaseName: pulumi.String("admin"),
			Roles:            roles,
		})
		if err != nil {
			return err
		}

		p.ClusterURL = pulumi.All(p.Cluster.ConnectionStrings.StandardSrv(), user.Username, user.Password).ApplyT(func(all []interface{}) string {
			username := all[1].(string)
			password := all[2].(string)

			connectionString := strings.Split(all[0].(string), "//")[1]

			return fmt.Sprintf("mongodb+srv://%s:%s@%s", username, password, connectionString)
		}).(pulumi.StringOutput)
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
