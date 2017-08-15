REGISTRY_SERVER=192.168.98.100:5000
NAME=greeting-api-doc
DOCKER_IMAGE=$(REGISTRY_SERVER)/$(NAME):latest

all: deploy_clean deploy_local

deploy_clean:
	kubectl delete deployment -l app=$(NAME)
	kubectl delete service -l app=$(NAME)

deploy_docker_kubernetes:
	docker build -t $(DOCKER_IMAGE) .
	docker push $(DOCKER_IMAGE)
	kubectl create -f deployment.yaml
deploy_local: deploy_docker_kubernetes


.PHONY:
	all \
	vet \
	build
