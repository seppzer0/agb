# build base image
FROM golang:1.23-alpine3.19 as base
ENV CGO_ENABLED=0
WORKDIR /build
COPY ./agb/* .
RUN go mod download && go generate ./...

# launch the linter
FROM golangci/golangci-lint:v1.60-alpine as lint
ENV CGO_ENABLED=0
WORKDIR /src
COPY --from=base /build .
COPY ./agb/.golangci.yaml .
RUN go mod download && golangci-lint run --timeout 5m

# build the binary
FROM base as build
WORKDIR /build
COPY --from=lint /src/lint_report.json .
RUN go test ./agb/... && go build -o /build/agb ./cmd/agb

# build final image
FROM alpine:3.19
WORKDIR /app
COPY --from=build /build/agb ./agb
ADD https://storage.googleapis.com/git-repo-downloads/repo /usr/local/bin/
RUN chmod -R 755 /usr/local/bin/repo

CMD [ "/bin/sh" ]
