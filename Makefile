.PHONY: build
build: vendor-proto .generate .build

PHONY: .generate
.generate:
		mkdir -p pkg/ocp-classroom-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-classroom-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-classroom-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-classroom-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-classroom-api \
				--swagger_out=allow_merge=true,merge_file_name=api:. \
				api/ocp-classroom-api/ocp-classroom-api.proto
		mv pkg/ocp-classroom-api/github.com/ozoncp/ocp-classroom-api/pkg/ocp-classroom-api/* pkg/ocp-classroom-api/
		rm -rf pkg/ocp-classroom-api/github.com
		mkdir -p cmd/ocp-classroom-api

PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=windows go build -o bin/ocp-classroom-api cmd/ocp-classroom-api/main.go

PHONY: install
install: build .install

PHONY: .install
install:
		go install cmd/grpc-server/main.go

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-classroom-api
		cp api/ocp-classroom-api/ocp-classroom-api.proto vendor.protogen/api/ocp-classroom-api
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
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get -u github.com/envoyproxy/protoc-gen-validate
