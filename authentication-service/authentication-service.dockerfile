FROM golang:1.18 as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build authApp .cmd/api
RUN chmod +x /app/brokerApp

FROM alplane:latest

RUN mkdir /app

COPY --from=builder /app/authApp /app

CMD [ "/app/authApp" ]