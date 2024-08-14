# syntax=docker/dockerfile:1

# Build Backend
FROM --platform=$BUILDPLATFORM golang:1.23.0-alpine3.19 as builder

LABEL org.opencontainers.image.title="Minetest Skin Server"
LABEL org.opencontainers.image.description="Skin server for the Minetest engine"
LABEL org.opencontainers.image.authors="AFCM <afcm.contact@gmail.com>"
LABEL org.opencontainers.image.licenses="GPL-3.0"
LABEL org.opencontainers.image.source="https://github.com/AFCMS/minetest-skin-server"

ARG TARGETOS
ARG TARGETARCH

ENV GOCACHE=/root/.cache/go-build

# Install build dependencies
RUN apk add --no-cache git make build-base
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN --mount=type=cache,id=gomod,target="/go/pkg/mod" go mod download
RUN --mount=type=cache,id=gomod,target="/go/pkg/mod" go mod verify

COPY . ./

# Build with cache
# https://dev.to/jacktt/20x-faster-golang-docker-builds-289n
RUN --mount=type=cache,id=gomod,target="/go/pkg/mod" \
    --mount=type=cache,id=gobuild,target="/root/.cache/go-build" \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o minetest-skin-server .

# Build Frontend
FROM --platform=$BUILDPLATFORM node:20-alpine3.19 as frontend-builder

WORKDIR /frontend

COPY ./frontend/package.json ./
COPY ./frontend/package-lock.json ./
RUN --mount=type=cache,id=npmmod,target="/root/.npm" npm ci

COPY ./frontend ./
RUN npm run build

# Production Image
FROM alpine:3.19 as production
RUN apk update && apk add --no-cache optipng

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "10001" \
    "appuser"

COPY --from=builder /app/minetest-skin-server /app/
COPY --from=builder /app/index.gohtml /app/
COPY --from=frontend-builder /frontend/dist /app/frontend/dist

USER appuser:appuser

WORKDIR /app

EXPOSE 8080
CMD ["/app/minetest-skin-server"]
