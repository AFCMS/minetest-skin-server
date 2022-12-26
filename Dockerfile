# Build Backend
FROM golang:1.19-alpine as builder

LABEL appuser="Minetest Skin Server"
LABEL author="AFCM <afcm.contact@gmail.com>"
LABEL description="Skin server for the Minetest engine"

RUN mkdir /build
COPY . /build
WORKDIR /build
RUN CGO_ENABLED=0 go build -o minetest-skin-server .

# Build Frontend
FROM node:16 as frontend-builder
RUN mkdir /build
COPY ./frontend ./build
WORKDIR /build
RUN npm install --include=dev && npm run build

# Production Image
FROM alpine:3.9.6 as production
RUN apk update && apk add --no-cache optipng=0.7.7-r1
COPY --from=builder /build/minetest-skin-server .
COPY --from=frontend-builder ./build ./frontend

EXPOSE 8080
CMD ["./minetest-skin-server"]