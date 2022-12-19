# Stage 1
FROM golang:1.19-alpine as builder

RUN apk update && apk add --no-cache git
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go get -d -v
RUN go build -o minetest-skin-server .

# Stage 2
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/ /app/
WORKDIR /app
CMD ["./minetest-skin-server"]