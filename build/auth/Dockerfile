FROM golang:1.22.2-alpine3.19 AS BUILD

WORKDIR /server

COPY . .

RUN go mod download

RUN go build -o /server/auth ./cmd/auth

FROM alpine:3.19

WORKDIR /server

COPY --from=BUILD /server/auth /server/auth

RUN chmod +x /server/auth

ENTRYPOINT [ "/server/auth" ]
