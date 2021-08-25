FROM golang:1.17-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /go/bin/webserver
COPY ./main.go .
RUN go mod init webserver
RUN go mod tidy
RUN GOOS=linux \
	GOARCH=amd64 \
	CGO_ENABLED=0 \
	go build -ldflags="-w -s" -o /bin/webserver

FROM scratch
LABEL app="webserver"
COPY --from=builder /bin/webserver /bin/webserver
EXPOSE 8080
ENTRYPOINT ["/bin/webserver"]
