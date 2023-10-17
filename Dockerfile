# Build Backend
FROM golang:1.21.3-alpine3.17 as builder

LABEL org.opencontainers.image.title="Minetest Skin Server"
LABEL org.opencontainers.image.description="Skin server for the Minetest engine"
LABEL org.opencontainers.image.authors="AFCM <afcm.contact@gmail.com>"
LABEL org.opencontainers.image.licenses="GPL-3.0"
LABEL org.opencontainers.image.source="https://github.com/AFCMS/minetest-skin-server"

RUN mkdir /build
COPY . /build
WORKDIR /build
RUN apk add --no-cache git make build-base
ENV CGO_ENABLED=1
RUN go build -o minetest-skin-server .

# Build Frontend
FROM node:18 as frontend-builder
RUN mkdir /build
COPY ./frontend /frontend
WORKDIR /frontend
RUN npm install --include=dev && npm run build

# Production Image
FROM alpine:3.18 as production
RUN apk update && apk add --no-cache optipng
COPY --from=builder /build/minetest-skin-server /
RUN mkdir -p /frontend/dist
COPY --from=frontend-builder /frontend/dist /frontend/dist

EXPOSE 8080
CMD ["./minetest-skin-server"]
