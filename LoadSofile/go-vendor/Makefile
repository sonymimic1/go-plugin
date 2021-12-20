#/bin/bash

NAME = $(shell basename $(shell pwd -P))
OUTPUT = $(NAME)
TARGET = ./main.go
LINUX_OS = linux
LINUX_ARCH = amd64


run-mod-vendor: 
	rm -rf ./vendor
	go mod vendor
	GO111MODULE=off GOOS=$(LINUX_OS) GOARCH=$(LINUX_ARCH) go run -gcflags='all=-N -l' main.go




