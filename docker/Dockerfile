# BUILD ENVIRONMENT
# -----------------
FROM golang:1.20-alpine as build_environment

WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download && apk add upx
COPY . .
RUN  go build -ldflags "-s -w -extldflags '-static'" -o searcher && upx ./searcher


# DEPLOYMENT ENVIRONMENT
# -----------------
FROM alpine:latest

RUN apk update && apk add --no-cache bash
WORKDIR /app
COPY --from=build_environment /src/searcher /app/

ENTRYPOINT ["./searcher"]