NAME = crypto-broker
VERSION=1.0.0
LOGGER_LEVEL=debug

config: # Copy config yaml
	cp config_example.yaml config.yaml

run: # Run project locally
	VERSION=$(VERSION) \
	LOGGER_LEVEL=$(LOGGER_LEVEL) \
	go run cmd/server/main.go

build-docker: # Build docker image
	DOCKER_BUILDKIT=1 \
	docker build  \
		--target=build \
		--tag=$(NAME) \
		--file=./build/Dockerfile .

image: build-docker # Create release docker image.
	DOCKER_BUILDKIT=1 \
	docker build  \
		--target=image \
		--tag=$(NAME) \
		--file=./build/Dockerfile .

run-docker: image # Run docker container.
	docker run --rm \
		--name $(NAME) \
		$(NAME)

build-local: ## Build binary locally
	-rm ./$(NAME)

	CGO_ENABLED=0 \
	GOOS=linux  \
	GOARCH=amd64  \
	go build -installsuffix cgo -o $(NAME) \
	./cmd/server/main.go

run-local: build-local ## Use it when you want to get your local enviroment up
	LOGGER_LEVEL=$(LOGGER_LEVEL) \
	VERSION=$(VERSION) \
	./$(NAME)