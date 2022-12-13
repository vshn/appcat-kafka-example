## These are some common variables for Make
PROJECT_ROOT_DIR = .
PROJECT_NAME ?= appcat-kafka-example
PROJECT_OWNER ?= vshn

WORK_DIR = $(PWD)/.work

## BUILD:go
BIN_FILENAME ?= $(PROJECT_NAME)
go_bin ?= $(WORK_DIR)/bin
$(go_bin):
	@mkdir -p $@

## BUILD:docker
DOCKER_CMD ?= docker

IMG_TAG ?= latest
# Image URL to use all building/pushing image targets
CONTAINER_IMG ?= ghcr.io/$(PROJECT_OWNER)/$(PROJECT_NAME):$(IMG_TAG)
