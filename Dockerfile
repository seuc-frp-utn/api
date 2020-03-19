# Builder
FROM golang:1.13.5 AS builder
LABEL maintainer="Secretaria de Extension Universitaria y Cultura <extension@frp.utn.edu.ar>"

## Enable go modules
ENV GO111MODULE=on \
    CGO_ENABLED=1

## Download dependencies
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

## Build application
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o api ./cmd/app

## Move binary file to dist
WORKDIR /dist
RUN cp /build/cmd/app/api .
RUN mkdir /data

#####################################################################

# Runner
FROM scratch
COPY --chown=0:0 --from=builder /dist /
COPY --chown=65534:0 --from=builder /data /data
USER 65534
WORKDIR /data
EXPOSE 3000
CMD ["/api"]