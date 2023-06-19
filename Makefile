APP = app
BUILD_DIR = build
CGO_ENABLED ?= 0
GOTAGS ?= musl
GOOS ?= linux
REGISTRY = registry.gitlab.com
DOCKERFILE_PATH ?= ./docker/Dockerfile

build:
	go build -mod=vendor -tags $(GOTAGS) -ldflags '-s -w -extldflags "-static"' -o ./${BUILD_DIR}/${APP} ./cmd/app/main.go