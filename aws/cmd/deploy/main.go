package main

import (
	_ "embed"

	"github.com/nitrictech/mongodb-provider/aws/deploy"
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
)

//go:embed runtime-extension-aws
var runtimeBin []byte

var runtimeProvider = func() []byte {
	return runtimeBin
}

// Start the deployment server
func main() {
	stack := deploy.NewAwsExtensionProvider()

	providerServer := provider.NewPulumiProviderServer(stack, runtimeProvider)

	providerServer.Start()
}
