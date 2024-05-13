FROM alpine:latest AS build
WORKDIR /app
COPY . .
RUN apk --update add gcc libc-dev
RUN gcc -o app -O3 main.c


FROM alpine:latest
WORKDIR /app
COPY --from=build /app/app .
CMD ["./app"]