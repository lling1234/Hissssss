.PHONY: build run upx test shbuild docker swag db buildx

build:
	@echo "Build"
	go build -o=bin/account account.go

run:
	@echo "Build and Run"
	CGO_ENABLED=0 go build -o=bin/account account.go && \
	cd bin && ./account
