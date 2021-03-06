FROM ubuntu
WORKDIR /app

EXPOSE 50050

ENV GRPC_PORT 50050
ENV GRPC_MODULES host
ENV GRPC_SERVICE_NAME host

RUN apt update && apt install -y curl unzip git wget

COPY ["golangServer.go", "golangServer/"]
COPY modules/ golangServer/modules/
COPY proto/ golangServer/proto/
COPY BuildFiles/ golangServer/BuildFiles/
COPY Config/ golangServer/Config/

WORKDIR golangServer/BuildFiles/

RUN wget https://golang.org/dl/go1.15.3.linux-amd64.tar.gz
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/protoc-3.13.0-linux-x86_64.zip

RUN tar -C /usr/local -xzf go1.15.3.linux-amd64.tar.gz
RUN unzip protoc-3.13.0-linux-x86_64.zip -d /root/.local

ENV PATH=$PATH:/usr/local/go/bin
ENV PATH="$PATH:/root/.local/bin"

RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go get github.com/BurntSushi/toml


ENV PATH="$PATH:/root/go/bin"

WORKDIR /app/golangServer/

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
