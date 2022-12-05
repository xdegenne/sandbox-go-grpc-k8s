IMAGE_NAME=sandbox-hello
IMAGE_VERSION=0.0.8


format:
	go fmt ./...
	go vet ./...

build:
	go build .

lint:
	golint ./...

image:
	docker build -t $(IMAGE_NAME):$(IMAGE_VERSION) -t $(IMAGE_NAME):latest .

push:
	docker push $(IMAGE_NAME):$(IMAGE_VERSION)
