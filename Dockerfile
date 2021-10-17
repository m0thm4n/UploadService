FROM golang:1.17.1-buster as base

FROM base as dev

RUN go version

RUN apt update
RUN apt install libexif-dev -y

ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api
CMD ["air"]

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 30/tcp

RUN go build -o /docker-upload-service
