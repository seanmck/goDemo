FROM golang:1.13.8-alpine3.10
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o main .
CMD ["/app/main"]
