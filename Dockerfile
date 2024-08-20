FROM golang:1.22.2 AS builder

WORKDIR /travel-api

COPY . .
RUN go mod download

COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../travel_app .

FROM alpine:latest

WORKDIR /travel-api 

COPY --from=builder /travel-api/travel_app .
COPY --from=builder /travel-api/logs/app.log ./logs/
COPY --from=builder /travel-api/.env .

EXPOSE 8080

CMD [ "./travel_app" ]