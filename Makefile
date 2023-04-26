VERSION=latest

.PHONY: build
build:
	go generate ./config/var.go
	GOOS=linux GOARCH=amd64 go build -o ./build/oracle-go main.go
	docker buildx build --platform=linux/amd64 -t nodedao/oracle-go:$(VERSION) -f ./build/Dockerfile-local .


.PHONY: push
push:
	docker push nodedao/oracle-go:$(VERSION)

