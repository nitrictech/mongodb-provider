package main

import (
	_ "embed"

	"github.com/nitrictech/mongodb-provider/azure/deploy"
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
)

//go:embed runtime-extension-azure
var runtimeBin []byte

var runtimeProvider = func() []byte {
	return runtimeBin
}

// Start the deployment server
func main() {
	stack := deploy.NewAzureExtensionProvider()

	providerServer := provider.NewPulumiProviderServer(stack, runtimeProvider)

	providerServer.Start()
}
