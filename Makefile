.DEFAULT_GOAL := docker-image

IMAGE ?= registry.cn-hangzhou.aliyuncs.com/linkedcare/tags-for-ecs:latest

bin/tags-for-ecs: $(shell find . -name '*.go')
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o $@ ./cmd/tags-for-ecs

.PHONY: docker-image
docker-image: bin/tags-for-ecs
	docker build -t $(IMAGE) --file image/Dockerfile .

.PHONY: push-image
push-image: docker-image
	docker push $(IMAGE)
