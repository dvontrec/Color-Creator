FROM golang:latest

WORKDIR "/app"

COPY . .

RUN go get -d -v ./...; go build -o server
ENTRYPOINT ["./server"]

