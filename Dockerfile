# Build Backend
FROM golang:1.19-alpine as builder

LABEL appuser="Minetest Skin Server"
LABEL author="AFCM <afcm.contact@gmail.com>"
LABEL description="Skin server for the Minetest engine"

RUN apk update && apk add --no-cache git optipng gcc && \
    mkdir /build
COPY . /build/
WORKDIR /build
RUN go get -d -v && \
    go build -o minetest-skin-server .

# Stage 2
FROM alpine:latest
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/ /app/
WORKDIR /app
CMD ["./minetest-skin-server"]