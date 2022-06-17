FROM golang:1.18-bullseye as builder

WORKDIR /app/

COPY ./src/server .

RUN ls -alh

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o /grpc

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/ /app/

CMD ./grpc

