## gRPC Example Version 1

> This version show a very simple version of gRPC
>
> - Define a simple Greeting service with gRPC and write a client to test it

## Technologies used in this example:

### Server technologies:

- Google Protocol Buffer
- gRPC
- gRPC related libs: protoc-gen-go

## Prerequisite: 

- Protocol buffer is installed

## How to run & test:

```sh
# Build proto
cd proto
./genproto

# Build and run the server
cd ../server
go build -o server.bin
./server.bin &

# Run the client to test the server
cd ../client
go run main.go

```

