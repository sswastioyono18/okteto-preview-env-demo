FROM golang:1.15.7-buster as build-env
ARG APPLICATION_NAME=oktetotest
WORKDIR /opt/$APPLICATION_NAME
ENV PATH=$PATH:/opt/$APPLICATION_NAME
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN --mount=type=cache,id=gomod,target=/go/pkg/mod \
    --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
      go build -o bin/${APPLICATION_NAME}

FROM alpine as alpine
RUN apk --no-cache add tzdata ca-certificates

FROM bgruening/busybox-bash
ARG APPLICATION_NAME=oktetotest
WORKDIR /opt/$APPLICATION_NAME
COPY --from=build-env /opt/$APPLICATION_NAME/bin/$APPLICATION_NAME /opt/$APPLICATION_NAME/$APPLICATION_NAME
COPY --from=alpine /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Jakarta
USER root

ENTRYPOINT "./oktetotest"

