PROTODIR := .
PROTOC := protoc
PROTOC_INCLUDES := /usr/local/include
PROTO_OUT = $(realpath ${GOPATH}/src)
API_DIR ?= $(PROTO_OUT)/github.com/golovers/grpc/v2/api-golang/gr
DOC_DIR ?= $(PROTO_OUT)/github.com/golovers/grpc/v2/api-doc

GRPC_GATEWAY=vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

all: proto

proto: init proto_go proto_gateway proto_swagger

init:
	cd vendor/github.com/golang/protobuf/protoc-gen-go && go build && go install
	cd vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && go build && go install
	cd vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && go build && go install

proto_go:
	@ if ! which $(PROTOC) > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	for dir in $$(find . -name '*.proto' -exec dirname '{}' ';' |  grep gr/ | sort -u); do \
		protoc -I gr/ -I $(PROTOC_INCLUDES) -I $(GRPC_GATEWAY) $$dir/*.proto --go_out=plugins=grpc:$(API_DIR) || exit 1; \
	done
proto_gateway:
	@ if ! which $(PROTOC) > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	for dir in $$(find . -name '*.proto' -exec dirname '{}' ';' |  grep gr/ | sort -u); do \
		protoc -I gr/ -I $(PROTOC_INCLUDES) -I $(GRPC_GATEWAY) $$dir/*.proto --grpc-gateway_out=logtostderr=true:$(API_DIR) || exit 1; \
	done
proto_swagger:
	@ if ! which $(PROTOC) > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	for dir in $$(find . -name '*.proto' -exec dirname '{}' ';' |  grep gr/ | sort -u); do \
		protoc -I gr/ -I $(PROTOC_INCLUDES) -I $(GRPC_GATEWAY) $$dir/*.proto --swagger_out=logtostderr=true:$(DOC_DIR) || exit 1; \
	done
.PHONY:
	all \
	proto \
	proto_go \
	clean \
	gen
