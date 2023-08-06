.PHONY: build run upx test shbuild docker swag db buildx

build:
	@echo "Build"
	CGO_ENABLED=0 go build -o=bin/job job.go

run:
	@echo "Build and Run"
	CGO_ENABLED=0 go build -o=bin/job job.go && \
	cd bin && ./job
