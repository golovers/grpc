NAME=greeting

all: vet build

vet:
	go vet .

build:
	CGO_ENABLED=0 go build -tags netgo -o $(NAME).bin

format:
	gofmt -l -w .


.PHONY:
	all \
	vet \
	build
