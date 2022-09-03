FROM golang:1.18-alpine as BUILD
WORKDIR /go/src/github.com/demeero/shopagolic/productcatalog

# Explicitly copy go.mod and go.sum, then download the vendor dir
# allowing for this step to be cached more often during development.
COPY go.* ./
RUN go mod download

COPY . .
