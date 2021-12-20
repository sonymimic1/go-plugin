#/bin/bash

NAME = $(shell basename $(shell pwd -P))
OUTPUT = $(NAME)
TARGET = ./main.go
LINUX_OS = linux
LINUX_ARCH = amd64

run-mod: 
	GO111MODULE=on go run -trimpath -gcflags='all=-N -l' main.go




