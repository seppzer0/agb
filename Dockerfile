# build base image
FROM golang:1.24-alpine3.22 AS base
ENV CGO_ENABLED=0
WORKDIR /build
COPY ./agb .
RUN go mod download && go generate ./...

# launch the linter
FROM golangci/golangci-lint:v2.1-alpine AS lint
ENV CGO_ENABLED=0
WORKDIR /src
COPY --from=base /build .
COPY ./agb/.golangci.yaml .
RUN go mod download && golangci-lint run --timeout 5m

# build the binary
FROM base AS build
WORKDIR /build
COPY --from=lint /src/lint_report.json .
RUN ls .
RUN go test ./... && go build -o /build/agb ./cmd/agb

# build final image
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /build/agb ./agb
ADD https://storage.googleapis.com/git-repo-downloads/repo /usr/local/bin/
RUN chmod -R 755 /usr/local/bin/repo
RUN \
    apt-get update \
    && \
    apt-get install -y \
        curl \
        wget \
        git \
        gcc \
        g++ \
        libssl-dev \
        python3 \
        python3-pip \
        make \
        zip \
        bc \
        libgpgme-dev \
        bison \
        flex

CMD [ "/bin/sh" ]
