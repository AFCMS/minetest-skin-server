# Build Backend
FROM golang:1.19-alpine3.17 as builder

LABEL appuser="Minetest Skin Server"
LABEL author="AFCM <afcm.contact@gmail.com>"
LABEL description="Skin server for the Minetest engine"

RUN mkdir /build
COPY . /build
WORKDIR /build
RUN apk add --no-cache git=2.38.2-r0 make=4.3-r1 build-base=0.5-r3
ENV CGO_ENABLED=1
RUN go build -o minetest-skin-server .

# Build Frontend
FROM node:16 as frontend-builder
RUN mkdir /build
COPY ./frontend /frontend
WORKDIR /frontend
RUN npm install --include=dev && npm run build

# Production Image
FROM alpine:3.17 as production
RUN apk update && apk add --no-cache optipng=0.7.7-r1
COPY --from=builder /build/minetest-skin-server /
RUN mkdir -p /frontend/build
COPY --from=frontend-builder /frontend/build /frontend/build

EXPOSE 8080
CMD ["./minetest-skin-server"]