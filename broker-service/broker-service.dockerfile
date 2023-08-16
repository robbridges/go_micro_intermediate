#base go image
FROM golang:1.20-alpine as builder

RUN mkdir /app

COPY brokerApp /app

CMD [ "/app/brokerApp"]
