FROM golang:1.19.0

# enter own pass to project files
WORKDIR /go/gogin

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy