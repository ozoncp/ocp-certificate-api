module github.com/ozoncp/ocp-certificate-api/test/load

go 1.16

require (
	github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api v0.0.0-20210823131530-e588ca2f87ad
	github.com/rs/zerolog v1.23.0
	github.com/spf13/afero v1.6.0
	github.com/yandex/pandora v0.3.3
	go.uber.org/zap v1.19.0 // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api => ./../../pkg/ocp-certificate-api
