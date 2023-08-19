FROM golang:1.20-alpine


RUN mkdir app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./
ENV GO111MODULE=on
ENV GOOS=linux
ENV CGO_ENABLED=0
EXPOSE 8080

ENTRYPOINT go run .

