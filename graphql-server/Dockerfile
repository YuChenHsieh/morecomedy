#first stage - builder
FROM golang:1.17.0-stretch as builder
COPY . /graphql
WORKDIR /graphql
ENV GO117MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/graphql-server/main.go 
#second stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /graphql/app .
CMD ["./app"]