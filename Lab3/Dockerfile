FROM golang:latest

WORKDIR /Lab3

COPY . .

RUN go mod tidy
RUN go build -o laboratorio Tierra/tierra.go

EXPOSE 50051

CMD ["./laboratorio"]
