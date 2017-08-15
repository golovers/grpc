## gRPC Example Version 3

> This version show a very simple version of gRPC
>
> - Define a simple Greeting service with both gRPC and REST API exposed
> - Generate Swagger doc with go protoc-gen-swagger
> - Deploy the components with Docker and Kubernetes

## Technologies used in this example:

### Server technologies:

- Google Protocol Buffer
- gRPC
- gRPC related libs: protoc-gen-go, protoc-gen-swagger, protoc-gen-grpc-gateway

### Build  & Deployment technologies

- [Glide](https://glide.sh/)
- [Docker](https://www.docker.com/)
- Kubernetes ([Minikube](https://github.com/kubernetes/minikube))

## Prerequisite: 

- Glide is installed
- Protocol buffer is installed
- Docker private registry and  Kubernetes are installed and started

## How to run:

```sh
# Build protobuf files
cd apis
glide install
make

# Build the server
cd ../greeting
glide install
make

# Build the client and run the test
cd ../client
glide install
go run main.go

# Build the doc and deploy with Swagger UI
cd ../api-doc
make
minikube service greeting-api-doc
```

