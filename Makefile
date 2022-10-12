SHELL:= /bin/bash

.DEFAULT_GOAL := init

.PHONY: init
init: init.dir 

.PHONY: build
build: api.server comment.server

ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(shell pwd)
endif

ifeq ($(origin DATA_DIR),undefined)
DATA_DIR := $(ROOT_DIR)/data
endif


ifeq ($(origin API_SRV_DIR),undefined)
USER_SRV_DIR := $(ROOT_DIR)/cmd/client
endif


ifeq ($(origin PIR_SRV_DIR),undefined)
USER_SRV_DIR := $(ROOT_DIR)/cmd/server
endif


.PHONY: init.dir
init.dir:
	@echo "===========> Run init.dir"
	@echo "ROOT_DIR: "$(ROOT_DIR) 
	@echo "DATA_DIR: "$(DATA_DIR) 
	

.PHONY: client
client:
	@echo "===========> Run api.server" 
	@cd $(API_SRV_DIR);go run client.go

.PHONY: server
server:
	@echo "===========> Run user.server" 
	@cd $(USER_SRV_DIR);go run server.go
