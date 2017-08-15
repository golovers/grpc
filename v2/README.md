## gRPC Example Version 2

> This version show a very simple version of gRPC
>
> - Define a simple Greeting service with both gRPC and REST API exposed
> - Generate Swagger doc with go protoc-gen-swagger

## How to run:

```sh
cd apis
glide install
make

cd ../greeting
glide install
make
./greeting.bin &

cd ../client
glide install
go run main.go
```

