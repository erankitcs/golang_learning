FROM golang:1.17-alpine as build

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build

ADD . /build/
WORKDIR /build
ADD server /build/

WORKDIR /build/server
RUN go mod tidy

RUN mkdir -p /dist

RUN go build -o /dist .

##
## Deploy
##

FROM alpine:latest
ENV tls=""
WORKDIR /cert
ADD cert .
WORKDIR /app
COPY --from=build /dist .

EXPOSE 9000
RUN echo "Running GRPC Server with mTLS as $mtls"
CMD ./server -"$tls"