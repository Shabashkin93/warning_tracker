FROM golang:1.21.4-alpine as builder

RUN apk add git
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /go/src/github.com/Shabashkin93/warning_tracker
COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . /go/src/github.com/Shabashkin93/warning_tracker
RUN CGO_ENABLED=0 GOOS=linux go generate ./...
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-X 'github.com/Shabashkin93/warning_tracker/pkg/buildinfo.version=v0.0.1' \
    -X 'github.com/Shabashkin93/warning_tracker/pkg/buildinfo.buildTime=$(date +”%Y.%m.%d.%H%M%S”)' \
    -X 'github.com/Shabashkin93/warning_tracker/pkg/buildinfo.commitHash=$(git rev-parse --short HEAD)'" main.go


FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/github.com/Shabashkin93/warning_tracker/main /usr/bin/github.com/Shabashkin93/warning_tracker

EXPOSE ${SERVER_PORT} ${SERVER_PORT}

ENTRYPOINT ["/usr/bin/github.com/Shabashkin93/warning_tracker"]
