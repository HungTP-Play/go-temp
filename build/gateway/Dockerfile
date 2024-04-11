FROM golang:1.22.2-alpine3.19 AS BUILD

WORKDIR /server

COPY . .

RUN go mod download

RUN go build -o /server/gateway ./cmd/gateway

FROM alpine:3.19

WORKDIR /server

COPY --from=BUILD /server/gateway /server/gateway

RUN chmod +x /server/gateway

ENTRYPOINT [ "/server/gateway" ]
