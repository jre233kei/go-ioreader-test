FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY usecase/*.go ./usecase/

RUN go install github.com/golang/mock/mockgen@v1.6.0  
RUN go generate -v ./...

CMD go test -v ./...