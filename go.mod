module github.com/nitrictech/mongodb-provider

go 1.21.4

replace (
	github.com/nitrictech/nitric/cloud/aws v0.0.0-20240510025749-b69ea254d49a => ../nitric/cloud/aws
	github.com/nitrictech/nitric/cloud/azure v0.0.0-20240510025749-b69ea254d49a => ../nitric/cloud/azure
	github.com/nitrictech/nitric/cloud/common v0.0.0-20240510025749-b69ea254d49a => ../nitric/cloud/common
	github.com/nitrictech/nitric/cloud/gcp v0.0.0-20240510025749-b69ea254d49a => ../nitric/cloud/gcp
	github.com/nitrictech/nitric/core v0.0.0-20240510025749-b69ea254d49a => ../nitric/core
)

require (
	github.com/charmbracelet/log v0.2.4
	github.com/mitchellh/mapstructure v1.5.0
	github.com/nitrictech/nitric/cloud/aws v0.0.0-20240510025749-b69ea254d49a
	github.com/nitrictech/nitric/cloud/azure v0.0.0-20240510025749-b69ea254d49a
	github.com/nitrictech/nitric/cloud/common v0.0.0-20240510025749-b69ea254d49a
	github.com/nitrictech/nitric/cloud/gcp v0.0.0-20240510025749-b69ea254d49a
	github.com/nitrictech/nitric/core v0.0.0-20240510025749-b69ea254d49a
	github.com/pulumi/pulumi-mongodbatlas/sdk/v2 v2.1.1
	github.com/pulumi/pulumi-random/sdk/v4 v4.8.2
	github.com/pulumi/pulumi/sdk/v3 v3.105.0
	github.com/samber/lo v1.38.1
	go.mongodb.org/mongo-driver v1.15.0
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
)

require (
	cloud.google.com/go v0.112.2 // indirect
	cloud.google.com/go/apigateway v1.6.6 // indirect
	cloud.google.com/go/cloudtasks v1.12.7 // indirect
	cloud.google.com/go/compute v1.25.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/firestore v1.15.0 // indirect
	cloud.google.com/go/iam v1.1.7 // indirect
	cloud.google.com/go/longrunning v0.5.6 // indirect
	cloud.google.com/go/pubsub v1.37.0 // indirect
	cloud.google.com/go/secretmanager v1.12.0 // indirect
	cloud.google.com/go/storage v1.39.1 // indirect
	dario.cat/mergo v1.0.0 // indirect
	github.com/Azure/azure-pipeline-go v0.2.3 // indirect
	github.com/Azure/azure-sdk-for-go v67.1.0+incompatible // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.9.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.5.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/data/aztables v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.5.1 // indirect
	github.com/Azure/azure-storage-blob-go v0.15.0 // indirect
	github.com/Azure/azure-storage-queue-go v0.0.0-20230927153703-648530c9aaf2 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.28 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.21 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.11 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.6 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.2.1 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator v0.34.2 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230828082145-3c4c8a2d2371 // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/avast/retry-go v3.0.0+incompatible // indirect
	github.com/aws/aws-lambda-go v1.34.1 // indirect
	github.com/aws/aws-sdk-go v1.44.298 // indirect
	github.com/aws/aws-sdk-go-v2 v1.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.1 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.27.4 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.4 // indirect
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.13.6 // indirect
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression v1.7.6 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.15.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi v1.19.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.20.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.20.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.9.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.17.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.21.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.51.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.28.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sfn v1.26.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sns v1.29.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.31.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.23.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.1 // indirect
	github.com/aws/smithy-go v1.20.1 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/charmbracelet/bubbles v0.16.1 // indirect
	github.com/charmbracelet/bubbletea v0.24.2 // indirect
	github.com/charmbracelet/lipgloss v0.8.0 // indirect
	github.com/cheggaaa/pb v1.0.29 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/containerd/console v1.0.4-0.20230313162750-1ae8d489ac81 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/djherbis/times v1.5.0 // indirect
	github.com/docker/distribution v2.8.0+incompatible // indirect
	github.com/docker/docker v20.10.24+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fasthttp/router v1.4.18 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/getkin/kin-openapi v0.113.0 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.5.0 // indirect
	github.com/go-git/go-git/v5 v5.11.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.0 // indirect
	github.com/golang/glog v1.2.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.3 // indirect
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/hcl/v2 v2.17.0 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/invopop/yaml v0.1.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-ieproxy v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/montanaflynn/stats v0.7.0 // indirect
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/nitrictech/nitric/test v0.0.0-20240510025749-b69ea254d49a // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opentracing/basictracer-go v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/perimeterx/marshmallow v1.1.4 // indirect
	github.com/pgavlin/fx v0.1.6 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/term v1.1.0 // indirect
	github.com/pulumi/appdash v0.0.0-20231130102222-75f619a67231 // indirect
	github.com/pulumi/esc v0.6.2 // indirect
	github.com/pulumi/pulumi-aws/sdk/v5 v5.26.0 // indirect
	github.com/pulumi/pulumi-azure-native-sdk v1.101.0 // indirect
	github.com/pulumi/pulumi-azure-native-sdk/apimanagement v1.92.0 // indirect
	github.com/pulumi/pulumi-azure-native-sdk/app v1.92.0 // indirect
	github.com/pulumi/pulumi-azure-native-sdk/authorization v1.92.0 // indirect
	github.com/pulumi/pulumi-azure-native-sdk/containerregistry v1.92.0 // indirect
	github.com/pulumi/pulumi-azure-native-sdk/eventgrid v1.101.0 // indirect
	github.com/pulumi/pulumi-azure-native-sdk/keyvault v1.92.0 // indirect
	github.com/pulumi/pulumi-azure-native-sdk/managedidentity v1.92.0 // indirect
	github.com/pulumi/pulumi-azure-native-sdk/operationalinsights v1.92.0 // indirect
	github.com/pulumi/pulumi-azure-native-sdk/resources v1.92.0 // indirect
	github.com/pulumi/pulumi-azure-native-sdk/storage v1.92.0 // indirect
	github.com/pulumi/pulumi-azure/sdk/v4 v4.42.0 // indirect
	github.com/pulumi/pulumi-azuread/sdk/v5 v5.33.0 // indirect
	github.com/pulumi/pulumi-docker/sdk/v4 v4.1.0 // indirect
	github.com/pulumi/pulumi-gcp/sdk/v6 v6.55.0 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06 // indirect
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.0 // indirect
	github.com/savsgio/gotils v0.0.0-20230208104028-c358bd845dee // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/skeema/knownhosts v1.2.1 // indirect
	github.com/spf13/cobra v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/texttheater/golang-levenshtein v1.0.1 // indirect
	github.com/tweekmonster/luser v0.0.0-20161003172636-3fa38070dbd7 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.45.0 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	github.com/zclconf/go-cty v1.13.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.36.4 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.49.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.49.0 // indirect
	go.opentelemetry.io/contrib/propagators/aws v1.11.0 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/sdk v1.22.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/exp v0.0.0-20240103183307-be819d1f06fc // indirect
	golang.org/x/mod v0.15.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/oauth2 v0.18.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/term v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	google.golang.org/api v0.170.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240325203815-454cdb8f5daa // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240325203815-454cdb8f5daa // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	lukechampine.com/frand v1.4.2 // indirect
)
