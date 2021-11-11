FROM golang:1.16-alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN GO111MODULE=on go mod download
#WORKDIR /go/src/app/cmd/tss
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/src


FROM alpine:latest
ARG privkey
ARG net
ENV PRIVKEY=$privkey
ENV NET=$net
RUN apk add --update ca-certificates curl
RUN mkdir -p /go/bin
COPY --from=builder /go/bin/src .
COPY build/start-tss.bash /go/bin/start-tss.bash
COPY build/start.bash /go/bin/start.bash
EXPOSE 6668
EXPOSE 8080
RUN chmod +x /go/bin/start-tss.bash
RUN chmod +x /go/bin/start.bash