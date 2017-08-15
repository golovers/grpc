## gRPC Example Version 4

> This version show a very simple version of gRPC
>
> - Define a simple Greeting service with both gRPC and REST API exposed
> - Generate Swagger doc with go protoc-gen-swagger
> - Deploy the components with Docker and Kubernetes
> - **Demonstrate how to use gRPC middleware (in this example is monitor gRPC with Prometheus)**
> - **Enable TLS for both gRPC and REST API of the Greeting service**

## Technologies used in this example:

### Server technologies:

- Google Protocol Buffer
- gRPC
- gRPC related libs: protoc-gen-go, protoc-gen-swagger, protoc-gen-grpc-gateway, go-grpc-prometheus,  go-grpc-middleware (see details [here](github.com/grpc-ecosystem))
- Prometheus

### Build  & Deployment technologies

- [Glide](https://glide.sh/)
- [Docker](https://www.docker.com/)
- Kubernetes ([Minikube](https://github.com/kubernetes/minikube))

## Prerequisite: 

- Glide is installed
- Protocol buffer is installed
- Docker private registry and  Kubernetes are installed and started

## How to run & test:

```sh
# Build protobuf files
cd apis
glide install
make

# Build the server and deploy to Kubernetes
cd ../greeting
glide install
make

# Build the client and run the test on both gRPC and REST API
cd ../client
glide install
go run main.go

# Build the doc, deploy to Kubernetes and Open browser for testing Swagger UI
cd ../api-doc
make
minikube service greeting-api-doc
cd ..

# Get to see statistic of gRPC metrics using Prometheus
curl 192.168.99.100:30104/metrics
```

