FROM golang:latest

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o auth .

EXPOSE 5001

CMD ["./auth"]
