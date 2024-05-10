package main

import (
	"os"
	"os/signal"
	"syscall"

	mongo_service "github.com/nitrictech/mongodb-provider/common"
	"github.com/nitrictech/nitric/cloud/aws/runtime/api"
	"github.com/nitrictech/nitric/cloud/aws/runtime/env"
	lambda_service "github.com/nitrictech/nitric/cloud/aws/runtime/gateway"
	sqs_service "github.com/nitrictech/nitric/cloud/aws/runtime/queue"
	"github.com/nitrictech/nitric/cloud/aws/runtime/resource"
	secrets_manager_secret_service "github.com/nitrictech/nitric/cloud/aws/runtime/secret"
	s3_service "github.com/nitrictech/nitric/cloud/aws/runtime/storage"
	sns_service "github.com/nitrictech/nitric/cloud/aws/runtime/topic"
	"github.com/nitrictech/nitric/cloud/aws/runtime/websocket"
	base_http "github.com/nitrictech/nitric/cloud/common/runtime/gateway"
	"github.com/nitrictech/nitric/core/pkg/logger"
	"github.com/nitrictech/nitric/core/pkg/membrane"
)

func main() {
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt, syscall.SIGINT)

	logger.SetLogLevel(logger.INFO)

	gatewayEnv := env.GATEWAY_ENVIRONMENT.String()

	membraneOpts := membrane.DefaultMembraneOptions()

	provider, err := resource.New()
	if err != nil {
		logger.Fatalf("could not create aws provider: %v", err)
		return
	}

	// Load the appropriate gateway based on the environment.
	switch gatewayEnv {
	case "lambda":
		membraneOpts.GatewayPlugin, _ = lambda_service.New(provider)
	default:
		membraneOpts.GatewayPlugin, _ = base_http.NewHttpGateway(nil)
	}

	membraneOpts.ApiPlugin = api.NewAwsApiGatewayProvider(provider)
	membraneOpts.SecretManagerPlugin, _ = secrets_manager_secret_service.New(provider)
	membraneOpts.KeyValuePlugin, err = mongo_service.New()
	if err != nil {
		logger.Fatalf("There was an error initializing the mongo server: %v", err)
	}

	membraneOpts.TopicsPlugin, _ = sns_service.New(provider)
	membraneOpts.StoragePlugin, _ = s3_service.New(provider)
	membraneOpts.ResourcesPlugin = provider
	membraneOpts.WebsocketPlugin, _ = websocket.NewAwsApiGatewayWebsocket(provider)
	membraneOpts.QueuesPlugin, _ = sqs_service.New(provider)

	m, err := membrane.New(membraneOpts)
	if err != nil {
		logger.Fatalf("There was an error initializing the membrane server: %v", err)
	}

	errChan := make(chan error)
	// Start the Membrane server
	go func(chan error) {
		errChan <- m.Start()
	}(errChan)

	select {
	case membraneError := <-errChan:
		logger.Errorf("Membrane Error: %v, exiting\n", membraneError)
	case sigTerm := <-term:
		logger.Infof("Received %v, exiting\n", sigTerm)
	}

	m.Stop()
}
