# syntax=docker/dockerfile:1.7
FROM --platform=$BUILDPLATFORM golang:1.27-alpine@sha256:8a5910f31396cd4d89662f56c68b3ae31d374308270a1c3bd96672ee5ed43414 AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src
COPY go.mod main.go ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -o /counter .

FROM scratch
COPY --from=build /counter /counter
ENTRYPOINT ["/counter"]
