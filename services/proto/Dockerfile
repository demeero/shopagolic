FROM alpine:3.16 as base

RUN apk add --no-cache curl go protoc git

# Install the protocol compiler plugins for Go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Install buf.build
RUN VERSION=1.6.0 && \
    curl -sSL https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m) \
        -o /usr/local/bin/buf && \
    chmod +x /usr/local/bin/buf

WORKDIR proto
COPY . .

FROM base as lint
RUN export PATH=$PATH:$(go env GOPATH)/bin && buf lint

FROM base as gen
RUN export PATH=$PATH:$(go env GOPATH)/bin && buf generate
RUN cd gen/go && go mod init github.com/demeero/shopagolic/proto/gen/go && go mod tidy