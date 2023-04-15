FROM golang:1.20-buster as builder
ENV GO111MODULE=on
ENV GOPATH=/go
RUN echo $GOPATH
WORKDIR /app
ADD . .
RUN GOOS=linux go install github.com/NodeDAO/oracle-go

# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10:debug-nonroot
WORKDIR /app
COPY --from=builder /go/bin/oracle-go .
ENTRYPOINT ["./oracle-go", "oracle" , "withdraw"]
