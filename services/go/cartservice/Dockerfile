FROM golang:1.18-alpine AS builder

WORKDIR /go/src/github.com/demeero/shopagolic/cart

# Explicitly copy go.mod and go.sum, then download the vendor dir
# allowing for this step to be cached more often during development.
COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /api -v ./cmd/api

FROM alpine:3.16 AS final
ENV GRPC_PORT=8080
COPY --from=builder /api /
EXPOSE ${GRPC_PORT}
CMD ["/api"]
