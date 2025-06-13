FROM golang:1.24.4-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

RUN go build -o main .

# EXPOSE application port
EXPOSE 8000

CMD ["/app/main"]