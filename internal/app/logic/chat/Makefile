.PHONY: build run upx test shbuild docker swag db buildx

build:
	@echo "Build"
	CGO_ENABLED=0 go build -o=bin/chat chat.go

run:
	@echo "Build and Run"
	CGO_ENABLED=0 go build -o=bin/chat chat.go && \
	cd bin && ./chat
