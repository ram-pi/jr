FROM golang:1.20.1-buster AS builder

ARG VERSION=dev
ARG USER=docker

WORKDIR /go/src/app
COPY . .
#RUN go build -o main -ldflags=-X=main.version=${VERSION} main.go 
RUN go mod download
RUN go build -v -ldflags="-X 'jr/cmd.Version=${VERSION}' -X 'jr/cmd.BuildUser=${USER}' -X 'jr/cmd.BuildTime=`shell date`'" -o build/jr jr.go

FROM debian:buster-slim
COPY --from=builder /go/src/app/build/jr /go/bin/jr
COPY templates /root/.jr/templates/
ENV PATH="/go/bin:${PATH}"
CMD ["jr"]