FROM golang:1.20-buster as builder
ENV GO111MODULE=on
ENV GOPATH=/go
ENV GOPROXY=https://goproxy.cn
RUN echo $GOPATH
WORKDIR /app
ADD . .
RUN GOOS=linux go build -o /go/bin/oracle-go  main.go

# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10:debug-nonroot
WORKDIR /app
COPY --from=builder /go/bin/oracle-go .
ENTRYPOINT ["./oracle-go", "oracle" , "withdraw"]
