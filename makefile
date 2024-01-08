.PHONY: k8s_up k8s_down go_api go_msg docker build_docker push_docker rmi_docker cluster_up pod_up service_up local_up down service_down pod_down cluster_down 

DOCKER_ACC = guughs
APP_NAME = gointensive23
VERSION = latest
POD_FILE = k8s/deployment.yaml
SERVICE_FILE = k8s/service.yml
SLEEP_TIME = 15

all: go_api

k8s_up: docker cluster_up pod_up service_up local_up
k8s_down: service_down pod_down cluster_down rmi_docker
docker: build_docker push_docker

go_api:
	@echo "-----------------------------------------------------------------------------------------------------------------------"
	@echo "Running local..."
	go run ./cmd/api/main.go


go_msg:
	@echo "-----------------------------------------------------------------------------------------------------------------------"
	@echo "Running local..."
	go run ./cmd/api/main.go

build_docker:
	@echo "Building..."
	docker build -t $(DOCKER_ACC)/$(APP_NAME):${VERSION} .

push_docker:
	@echo "Pushing..."
	docker push $(DOCKER_ACC)/$(APP_NAME):${VERSION}

rmi_docker:
	@echo "Removing image..."
	docker rmi $(DOCKER_ACC)/$(APP_NAME):${VERSION}

cluster_up:
	@echo "-----------------------------------------------------------------------------------------------------------------------"
	@echo "Creating cluster..."
	kind create cluster
	@echo "Cluster Info!"
	kubectl cluster-info --context kind-kind
	@sleep ${SLEEP_TIME}

pod_up:
	@echo "-----------------------------------------------------------------------------------------------------------------------"
	@echo "Creating pod..."
	kubectl apply -f ${POD_FILE}
	@echo "Pod created!"
	@sleep ${SLEEP_TIME}
	@echo "Pod Info!"
	kubectl get pods

service_up:
	@echo "-----------------------------------------------------------------------------------------------------------------------"
	@echo "Creating service..."
	kubectl apply -f ${SERVICE_FILE}
	@echo "Service Info!"
	kubectl get services

local_up:
	@echo "-----------------------------------------------------------------------------------------------------------------------"
	@echo "Uping local service..."
	kubectl port-forward svc/goapp-service 8888:8888

service_down:
	@echo "-----------------------------------------------------------------------------------------------------------------------"
	@echo "Deleting service..."
	kubectl delete -f ${SERVICE_FILE}
	@echo "Service deleted!"

pod_down:
	@echo "-----------------------------------------------------------------------------------------------------------------------"
	@echo "Deleting pod..."
	kubectl delete -f ${POD_FILE}
	@echo "Pod deleted!"

cluster_down:
	@echo "-----------------------------------------------------------------------------------------------------------------------"
	@echo "Deleting cluster..."
	kind delete cluster
	@echo "Cluster deleted!"
