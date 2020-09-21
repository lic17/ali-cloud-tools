.DEFAULT_GOAL := docker-image

IMAGE ?= registry.cn-hangzhou.aliyuncs.com/linkedcare/tag-disks-for-ecs:latest

image/tag-disks-for-ecs: $(shell find . -name '*.go')
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o $@ ./cmd/tag-disks-for-ecs

.PHONY: docker-image
docker-image: image/tag-disks-for-ecs
	docker build -t $(IMAGE) bin/

.PHONY: push-image
push-image: docker-image
	docker push $(IMAGE)
