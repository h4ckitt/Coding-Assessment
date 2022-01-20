FROM golang:alpine

RUN addgroup -S appgroup && adduser -S dharmy -G appgroup

WORKDIR /home/dharmy/app

COPY go.mod go.mod

RUN go mod download

COPY . .

RUN go build .

USER dharmy

CMD ["./assessment"]

