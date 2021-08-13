LOCAL_BIN:=$(CURDIR)/bin

run:
	go run cmd/ocp-certificate-api/main.go

lint:
	golint ./...

test:
	go test -v ./...

.PHONY: deploy
deploy: .compose-build .compose-up .migrate

.PHONY: .compose-build
.compose-build:
	docker-compose build

.PHONY: .compose-up
.compose-up:
	docker-compose up -d

.PHONY: .migrate
.migrate:
	migrate -path ./migrations -database 'postgres://postgres:postgres@127.0.0.1:5433/postgres?sslmode=disable' up

.PHONY: build
build: vendor-proto .generate .build

.PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-certificate-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-certificate-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-certificate-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-certificate-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-certificate-api \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/ocp-certificate-api/ocp-certificate-api.proto
		mv pkg/ocp-certificate-api/github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api/* pkg/ocp-certificate-api/
		rm -rf pkg/ocp-certificate-api/github.com
		mkdir -p cmd/ocp-certificate-api
		cd pkg/ocp-certificate-api && ls go.mod || go mod init github.com/ozoncp/ocp-certificate-api/pkg/ocp-certificate-api && go mod tidy

.PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-certificate-api cmd/ocp-certificate-api/main.go

.PHONY: install
install: build .install

.PHONY: .install
install:
		go install cmd/ocp-certificate-api/main.go

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-certificate-api
		cp api/ocp-certificate-api/ocp-certificate-api.proto vendor.protogen/api/ocp-certificate-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init github.com/ozoncp/ocp-certificate-api
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u google.golang.org/grpc
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get -u github.com/envoyproxy/protoc-gen-validate
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/envoyproxy/protoc-gen-validate