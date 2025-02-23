FROM golang:1.24-alpine

WORKDIR /go/src/northwind_api
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
EXPOSE 8001
CMD ["go", "run", "main.go"]
