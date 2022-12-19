# Build Backend
FROM golang:1.19-alpine as builder

RUN <<EOR
    apk update && apk add --no-cache git optipng gcc
    mkdir /build
EOR
COPY . /build/
WORKDIR /build
RUN <<EOR
    go get -d -v
    go build -o minetest-skin-server .
EOR

# Stage 2
FROM alpine:latest
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/ /app/
WORKDIR /app
CMD ["./minetest-skin-server"]