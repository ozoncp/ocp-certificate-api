module github.com/ozoncp/ocp-certificate-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Masterminds/squirrel v1.5.0
	github.com/Shopify/sarama v1.29.1
	github.com/envoyproxy/protoc-gen-validate v0.6.1 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/golang/mock v1.6.0
	github.com/golang/snappy v0.0.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/klauspost/compress v1.13.4 // indirect
	github.com/lib/pq v1.10.2
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api v0.0.0-00010101000000-000000000000
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pressly/goose/v3 v3.0.0
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.30.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/rs/zerolog v1.23.0
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.5
	google.golang.org/genproto v0.0.0-20210825212027-de86158e7fda // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api => ./pkg/ocp-certificate-api
