FROM golang:1.24.2-alpine3.21 AS build
WORKDIR /go/src/github.com/alfanzain/go-keyboard-chiper

COPY go.mod go.sum ./
RUN go mod download -x

COPY ./cmd ./cmd
COPY ./internal ./internal

WORKDIR /go/src/github.com/alfanzain/go-keyboard-chiper/cmd
RUN go build -o server

FROM alpine:3.21
RUN apk add ca-certificates tzdata

COPY --from=build /go/src/github.com/alfanzain/go-keyboard-chiper/cmd/server ./server

ENTRYPOINT [ "./server" ]
