FROM alpine:latest

RUN mkdir /app

COPY frontendApp /app
COPY cmd/web/templates /app/cmd/web/templates

WORKDIR /app

CMD ["./frontendApp"]
