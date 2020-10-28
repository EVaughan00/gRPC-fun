FROM ubuntu
WORKDIR /app

EXPOSE 50051

ENV GRPC_PORT 50051
ENV GRPC_MODULES minimega

RUN apt update && apt install -y curl unzip git

COPY ["golangServer.go", "golangServer/"]
COPY modules/ golangServer/modules/
COPY proto/ golangServer/proto/
COPY BuildFiles/ golangServer/BuildFiles/

RUN tar -C /usr/local -xzf golangServer/BuildFiles/go1.15.3.linux-amd64.tar.gz
RUN unzip golangServer/BuildFiles/protoc-3.13.0-linux-x86_64.zip -d /root/.local

ENV PATH=$PATH:/usr/local/go/bin
ENV PATH="$PATH:/root/.local/bin"

RUN /usr/local/go/bin/go get -u github.com/golang/protobuf/protoc-gen-go
RUN /usr/local/go/bin/go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

ENV PATH="$PATH:/root/go/bin"

WORKDIR golangServer/

ENV GO111MODULE=on
RUN go mod init grpc/golangServer

RUN BuildFiles/protobufCompile.sh

WORKDIR /app/golangServer/proto/
RUN go build

WORKDIR /app/golangServer/modules/
RUN go build

WORKDIR /app/golangServer

RUN go install grpc/golangServer

CMD golangServer

## docker run -d --rm --name golang-server -p 50051:50051 golang-grpc-server