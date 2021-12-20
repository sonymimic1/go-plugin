#/bin/bash

NAME = $(shell basename $(shell pwd -P))
OUTPUT = $(NAME)
TARGET = ./main.go
LINUX_OS = linux
LINUX_ARCH = amd64


docker-env:
	docker run -it -v $(pwd):/go/src --name plugintest golang:1.15.6 /bin/bash

build-mod: 
	GO111MODULE=on GOOS=$(LINUX_OS) GOARCH=$(LINUX_ARCH) go build -trimpath -buildmode=plugin -gcflags='all=-N -l' -o ./$(OUTPUT).so $(TARGET)

build-mod-vendor:
	rm -rf ./vendor
	go mod vendor
	GO111MODULE=off GOOS=$(LINUX_OS) GOARCH=$(LINUX_ARCH) go build -gcflags='all=-N -l'  -buildmode=plugin -o ./$(OUTPUT).so $(TARGET)




