# BUILD ENVIRONMENT
# -----------------
FROM golang:1.16-alpine as build_environment

WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download && apk add upx
COPY . .
RUN  go build -ldflags "-s -w -extldflags '-static'" -o api-test && upx ./api-test


# DEPLOYMENT ENVIRONMENT
# -----------------
FROM alpine

RUN apk update && apk add --no-cache bash
WORKDIR /app
COPY --from=build_environment /src/api-test /app/

ENTRYPOINT ["./api-test"]