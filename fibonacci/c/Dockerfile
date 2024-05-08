FROM alpine:latest

WORKDIR /usr/src/app

COPY . .

RUN gcc -o app main.c

CMD ["./app"]
