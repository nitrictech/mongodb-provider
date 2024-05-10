binaries: deploybin

# build runtime binary directly into the deploy director so it can be embedded directly into the deployment engine binary
runtimebin:
	@echo Building Extension Runtime Server
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/runtime-extension-azure -ldflags="-s -w -extldflags=-static" ./cmd/runtime

predeploybin: runtimebin
	@cp bin/runtime-extension-azure cmd/deploy/runtime-extension-azure

deploybin: predeploybin
	@echo Building Extension Deployment Server
	@CGO_ENABLED=0 go build -o bin/deploy-extension -ldflags="-s -w -extldflags=-static" -ldflags="-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore" ./cmd/deploy
	@rm cmd/deploy/runtime-extension-azure

.PHONY: install
install: deploybin
	@echo installing extension deployment server to ${HOME}/.nitric/providers/mongo/azure-0.0.1
	@mkdir -p ${HOME}/.nitric/providers/mongo
	@rm -f ${HOME}/.nitric/providers/mongo/azure-0.0.1
	@cp bin/deploy-extension ${HOME}/.nitric/providers/mongo/azure-0.0.1