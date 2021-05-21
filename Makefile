.PHONY: all

default: build-web build-backend

build-web:
	cd web && yarn && yarn build && cp -r dist ../internal/static

build-backend:
	#GOOS=linux GOARCH=amd64
	GOPROXY=https://goproxy.cn go build -ldflags="-w -s" -o gexport ./cmd/gexport/main.go