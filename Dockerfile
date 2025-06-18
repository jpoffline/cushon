# STEP 1
FROM golang:1.24.4

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon


COPY . .
#ENTRYPOINT CompileDaemon --build="go build cmd/main.go" --command=./main
