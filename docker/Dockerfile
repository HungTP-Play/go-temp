# STAGE 1 >---------< BUILD

FROM golang:1.22-alpine3.18 AS BUILD

WORKDIR /server

COPY . .

RUN go mod download
RUN go build -o main main.go


# STAGE 2 >---------< RUN

FROM alpine:3.18

WORKDIR /app

COPY --from=BUILD /server/main .

CMD ["./main"]

