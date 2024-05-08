FROM alpine:latest
RUN apk update && apk add perl perl-dev
WORKDIR /app
COPY . .
CMD ["perl", "main.pl"]