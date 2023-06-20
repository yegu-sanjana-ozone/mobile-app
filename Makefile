GO = go 
GOBUILD = $(GO) build
GOTEST = $(GO) test
DOCKER_IMAGE = mobileapp

test:
	$(GOTEST) -v ./...

swagger: 
	swag init --dir ./cmd --pd --parseDepth 10

docker-build:
	docker build -t $(DOCKER_IMAGE) .

installmockery:
	go install github.com/vektra/mockery/v2@v2.15.0

generatemocks:
    rm -rf mocks/* && \
    mockery --all --keeptree --inpackage
