module github.com/ozoncp/ocp-certificate-api

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.1 // indirect
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/lyft/protoc-gen-star v0.5.3 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.23.0
	github.com/spf13/afero v1.6.0 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.39.1
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
)

replace github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api => ./pkg/ocp-certificate-api
