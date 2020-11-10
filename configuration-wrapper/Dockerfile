FROM ubuntu
WORKDIR /app

EXPOSE 50050

ENV GRPC_PORT 50050
ENV GRPC_MODULES host
ENV GRPC_SERVICE_NAME host

RUN apt update && apt install -y curl unzip git wget

COPY ["wrapper.go", "wrapper/"]
COPY modules/ wrapper/modules/
COPY proto/ wrapper/proto/
COPY BuildFiles/ wrapper/BuildFiles/
COPY Config/ wrapper/Config/

WORKDIR wrapper/BuildFiles/

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

WORKDIR /app/wrapper/

ENV GO111MODULE=on
RUN go mod init grpc/wrapper

RUN BuildFiles/protobufCompile.sh

WORKDIR /app/wrapper/proto/
RUN go build

WORKDIR /app/wrapper/modules/
RUN go build

WORKDIR /app/wrapper

RUN go install grpc/wrapper

CMD wrapper

## docker run -d --rm --name golang-server -p 50051:50051 golang-grpc-server
