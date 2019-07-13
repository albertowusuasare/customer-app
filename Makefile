ARTIFACT_ID := customer-svc
.PHONY: all	
all: deps install test run

.PHONY: deps 
deps:
	@echo "Fetching dependencies for project..."
	go get -v ./...

.PHONY: build
build:
	@echo "Building application..."
	go build ./...

.PHONY: test
test:
	@echo "Running all tests..."
	go test -v ./...

.PHONY: install
install: 
	@echo "Installing application..."
	go install ./...

.PHONY: run	
run: 
	@echo "Running application..."
	$(ARTIFACT_ID)

.PHONY: clean
clean: 
	@echo "Removing built artifact $(ARTIFACT_ID)..."
	rm $(GOPATH)/bin/$(ARTIFACT_ID)

.PHONY: lint
lint: 
	./script/lint.sh

.PHONY: sanity-check
sanity-check: deps lint test

.PHONY: int-test
int-test:
	go test -v github.com/albertowusuasare/customer-app/cmd/customer-svc/test/integration

# Tagged Artifact build

# For the purposes of deployments it is sometimes necessary to tag a built artifact. 
# To do this in an automated way, we append the latest commit hash to the name of the executable
# For example if the executable is 'customer-service', and the latest 12 digit commit hash is '0f8e60cda45b'
# Then the tagged artifact is 'customer-service_0f8e60cda45b'.

TAG := $(shell git rev-parse --short=12 HEAD)
TAGGED_ARTIFACT := $(ARTIFACT_ID)_$(TAG)

.PHONY: all_tag
all_tag: build_tag run_tag
.PHONY: build_tag
build_tag:
	@echo "Building $(TAGGED_ARTIFACT) ..."
	go build -o $(TAGGED_ARTIFACT) ./cmd/customer-svc

.PHONY: run_tag
run_tag:
	@echo "Running $(TAGGED_ARTIFACT) ..."
	./$(TAGGED_ARTIFACT)

.PHONY: clean_tag
clean_tag: 
	@echo "Cleaning $(TAGGED_ARTIFACT) ..."
	rm $(TAGGED_ARTIFACT)

