package deploy

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type MongoDBConfig struct {
	OrgId string `mapstructure:"orgId"`
}

func ConfigFromAttributes(attributes map[string]interface{}) (*MongoDBConfig, error) {
	config := &MongoDBConfig{}
	err := mapstructure.Decode(attributes, config)
	if err != nil {
		return nil, err
	}

	if config.OrgId == "" {
		return nil, fmt.Errorf("invalid configuration: require an organisation id")
	}

	return config, nil
}
