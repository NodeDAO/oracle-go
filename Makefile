VERSION=latest

.PHONY: build-local
build-local:
	CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -o ./build/oracle-go main.go
	docker buildx build --platform=linux/amd64 -t nodedao/oracle-go:$(VERSION) -f ./build/Dockerfile-local ./build


.PHONY: push
push:
	docker push nodedao/oracle-go:$(VERSION)

.PHONY: go-bindata
go-bindata:
	go generate ./config/var.go
