FROM golang:1.17-alpine as build

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build

ADD . /build/
WORKDIR /build
ADD client /build/

WORKDIR /build/client
RUN go mod tidy

RUN mkdir -p /dist

RUN go build -o /dist .

##
## Deploy
##

FROM alpine:latest
ENV tls=tls
ENV opt=1
ENV serverhost=""
ENV serverport=9000

WORKDIR /cert
ADD cert .
WORKDIR /app
COPY --from=build /dist .
RUN echo "Running GRPC client with option $opt , tls as $tls and  mTLS as $mtls"
CMD ./client -o $opt -$tls -serverhost $serverhost -serverport $serverport