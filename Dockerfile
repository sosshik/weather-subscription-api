FROM golang:1.24.2-alpine AS builder

WORKDIR /usr/local/src/weather-subscription-api

COPY . .

RUN apk add --no-cache bash

RUN go mod download

RUN go build -o weather-subscription-api ./cmd/weather-subscription-api/main.go


FROM alpine AS runner

WORKDIR /weather-subscription-api

RUN apk add --no-cache bash

COPY --from=builder /usr/local/src/weather-subscription-api/weather-subscription-api .
COPY /migrations ./migrations

CMD ["./weather-subscription-api"]