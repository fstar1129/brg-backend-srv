# STEP 1 build executable binary
FROM golang:1.14-alpine AS builder
# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
ENV builddir app
COPY . $builddir
WORKDIR $builddir/src
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/src
# STEP 2 build a small image
FROM alpine
RUN mkdir /app
WORKDIR /app
RUN  apk --no-cache add ca-certificates
# Copy our static executable.
COPY --from=builder /go/bin/src .
COPY --from=builder /go/app/src/config.json .
EXPOSE 8000
# # Run the src binary.
CMD ["./src"]