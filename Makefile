GO := go
PKGS :=  $(shell $(GO) list ./...)
UNIT_TEST_PKGS := $(shell $(GO) list ./... | grep -v /test/integration)
INTEGRATION_TEST_PKGS := ./test/integration

ARTIFACT_ID := customer-svc
MAIN_PATH := ./cmd/$(ARTIFACT_ID)

all: format lint build test

format:
	@echo ">> Formatting project ..."
	@$(GO) fmt $(PKGS)

deps:
	@echo ">> Fetching project dependencies ..."
	@$(GO) get -v $(PKGS)

lint: deps 
	@echo ">> Linting project ..."
	./script/lint.sh

build:
	@echo ">> Building application ..."
	@$(GO) build $(MAIN_PATH)

test: deps
	@echo ">> Running all tests ..."
	@$(GO) test -v $(PKGS)

install: 
	@echo ">> Installing application ..."
	@$(GO) install $(PKGS)

run: build 
	@echo ">> Running application ..."
	./$(ARTIFACT_ID)

clean: 
	@echo ">> Removing built artifact $(ARTIFACT_ID) ..."
	rm $(ARTIFACT_ID)

unit-test: deps
	@echo ">> Running unit tests ..."
	@$(GO) test -v $(UNIT_TEST_PKGS)

int-test: deps
	@echo ">> Running integration tests ..."
	@$(GO) test -v $(INTEGRATION_TEST_PKGS)

module:
	@echo ">> Creating go module"
	./script/gomod.sh	
# Tagged Artifact build

# For the purposes of deployments it is sometimes necessary to tag a built artifact. 
# To do this in an automated way, we append the latest commit hash to the name of the executable
# For example if the executable is 'customer-service', and the latest 12 digit commit hash is '0f8e60cda45b'
# Then the tagged artifact is 'customer-service_0f8e60cda45b'.

TAG := $(shell git rev-parse --short=12 HEAD)
TAGGED_ARTIFACT := $(ARTIFACT_ID)_$(TAG)


all_tag: build_tag run_tag

build_tag:
	@echo "Building $(TAGGED_ARTIFACT) ..."
	go build -o $(TAGGED_ARTIFACT) ./cmd/customer-svc


run_tag:
	@echo "Running $(TAGGED_ARTIFACT) ..."
	./$(TAGGED_ARTIFACT)


clean_tag: 
	@echo "Cleaning $(TAGGED_ARTIFACT) ..."
	rm $(TAGGED_ARTIFACT)

## Docker deploy: builds, tags a google container registry (gcr) tag and deploys the image to gcr

docker_deploy: docker_build docker_tag docker_push

## Docker build: builds a docker image with tag 'ARTIFACT_ID'

docker_build:
	@echo "Building docker image ..."
	docker build --tag=$(ARTIFACT_ID) .


GCR_IMAGE_ID := gcr.io/onua-246719/$(ARTIFACT_ID):$(TAG)
docker_tag:
	@echo "Building docker gcr image ..."
	docker tag $(ARTIFACT_ID) $(GCR_IMAGE_ID)

docker_push:
	@echo "Pushing image to google container registry ..."
	gcloud auth configure-docker
	docker push $(GCR_IMAGE_ID)


docker_run:
	@echo "Running docker ..."
	docker run -p 5090:5090 $(GCR_IMAGE_ID)

.PHONY: all deps build test install run	clean lint sanity-check int-test all_tag build_tag run_tag clean_tag docker_deploy docker_build docker_tag docker_push docker_run