include .env
export $(shell sed 's/=.*//' .env)

# $(SERVER_ROLE) come from environments
#                can be [develop, production]
#======================================================
# HTTP Server
#======================================================
PROJECT=api-my-domain-com

APP_NAME=api-my-domain-com
APP_LOG_NAME=api.my-domain.com
APP_VERSION=v0.0.1

PROJECT_ROOT_PATH=../../../..
INSTALL_PATH=/usr/local/bin

BUILD_PLATFORM=linux
BIN_PATH=bin/$(APP_NAME)/linux
APP_BINARY=$(BIN_PATH)/$(APP_NAME)

SERVICE_NAME="api-my-domain-com"
SERVICE_PATH=/etc/systemd/system

#=========================================================================
#
#=========================================================================
.PHONY: help build

help: Makefile
	@echo "Выберите опцию сборки "$(BINARY_NAME)":"
	@sed -n 's/^##//p' $< | column -s ':' |  sed -e 's/^/ /'

## build-app: Build App
build-app:
	cd $(PROJECT_ROOT_PATH)/ && \
	CGO_ENABLED=0 GOOS=$(BUILD_PLATFORM) GOARCH=amd64 go build -ldflags="-s -w" -o $(APP_BINARY) -v src/cmd/$(APP_NAME)/main.go && \
	chmod +x $(APP_BINARY);

## install-app: Create Logs dir and Symlink to app at /usr/local/bin/
install-app:
	sudo mkdir -p /var/log/$(APP_LOG_NAME) && \
    sudo chmod 777 /var/log/$(APP_LOG_NAME) && \
    sudo ln -sf $(shell pwd)/$(PROJECT_ROOT_PATH)/$(APP_BINARY) $(INSTALL_PATH)/$(APP_NAME);

## uninstall-app: Remove Symlink to app from /usr/local/bin/
uninstall-app:
	sudo rm -f $(INSTALL_PATH)/$(APP_NAME);

## install-service: Install App Like system service
install-service:
	sudo systemctl enable $(shell pwd)/$(PROJECT_ROOT_PATH)/configs/$(SERVER_ROLE)/$(SERVICE_NAME).service && \
	sudo systemctl daemon-reload;

## uninstall-service: Uninstall App from system services
uninstall-service:
	sudo systemctl disable $(shell pwd)/$(PROJECT_ROOT_PATH)/configs/$(SERVER_ROLE)/$(SERVICE_NAME).service && \
	sudo systemctl daemon-reload;

## service-stop: Stop app service
service-stop:
	sudo systemctl stop $(SERVICE_NAME);

## service-start: Start app service
service-start:
	sudo systemctl start $(SERVICE_NAME);

## service-status: Get Status app service
service-status:
	sudo systemctl status $(SERVICE_NAME);

## service-restart: Restart app service
service-restart:
	sudo systemctl restart $(SERVICE_NAME);
