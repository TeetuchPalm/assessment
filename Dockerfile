FROM golang:1.19-alpine

# Set working directory
WORKDIR /go/src/target

COPY go.mod go.sum ./

#COPY go.sum ./

RUN go mod download

COPY *.go ./


# Run tests
CMD CGO_ENABLED=0 go test --tags=integration ./...
