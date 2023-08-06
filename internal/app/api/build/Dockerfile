FROM golang:1.18 AS Builder
LABEL stage=Build
WORKDIR /app/build/
ENV CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN cd app/ && CGO_ENABLED=0 go build -o app main.go

FROM alpine:3.14
LABEL stage=Run
WORKDIR /app/release/
COPY --from=Builder /app/build /app/release/
EXPOSE 8080
RUN apk update && \
	apk add --no-cache \
	vim && \
	curl && \
	ca-certificates && \
	bash && \
	tzdata && \
	ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
	echo Asia/Shanghai > /etc/timezone

ENTRYPOINT ["./bin/app"]