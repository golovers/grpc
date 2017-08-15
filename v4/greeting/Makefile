REGISTRY_SERVER=192.168.98.100:5000
NAME=greeting
DOCKER_IMAGE=$(REGISTRY_SERVER)/$(NAME):latest

all: vet build deploy_clean deploy_local

vet:
	go vet .

build:
	CGO_ENABLED=0 go build -tags netgo -o $(NAME).bin

format:
	gofmt -l -w .

deploy_clean:
	kubectl delete deployment -l app=$(NAME)
	kubectl delete service -l app=$(NAME)


deploy_docker_kubernetes:
	docker build -t $(DOCKER_IMAGE) .
	docker push $(DOCKER_IMAGE)
	kubectl create -f deployment.yaml
deploy_local: build deploy_docker_kubernetes

gen_certs:
	openssl genrsa -out certs/server.key 2048
	openssl req -new -x509 -sha256 -key certs/server.key -out certs/server.crt -days 3650
	# As we're using self-signed CA hence need to update the trust store by running
	# following commands
	sudo cp certs/server.crt /usr/local/share/ca-certificates/
	sudo update-ca-certificates

.PHONY:
	all \
	vet \
	build
