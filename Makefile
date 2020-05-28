.DEFAULT_GOAL := build

# Aliases
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get

BUILD_DIR := out
BUILD_EXE := octogen
BUILD_TARGET := ${BUILD_DIR}/${BUILD_EXE}

BASEDIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

# Optional stuff for demos/samples


build:
	@-mkdir -p ${BUILD_DIR}
	@-echo "BUILD: ${BUILD_TARGET}"
	$(GO_BUILD) -o $(BUILD_TARGET) -v cmd/${BUILD_EXE}/main.go

run: build
	./${BUILD_TARGET}

sample: build
	@./${BUILD_TARGET}

test:
	$(GO_TEST) -v ./...

clean:
	$(GO_CLEAN)
	rm -f $(BUILD_TARGET)
