SHELL := /bin/bash

all: image

dev:
	go install -v cmd/...

image:
	GOOS=linux GOARCH=amd64 go build -o ./docker/checker .
	cd docker && docker build -t host-checker:latest . && cd -
	docker tag host-checker:latest reg.qiniu.com/exiaohao/host-checker:dev
	docker push reg.qiniu.com/exiaohao/host-checker:dev