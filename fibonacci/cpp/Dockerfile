FROM alpine:latest AS build
WORKDIR /app
COPY . .
RUN apk --update add g++
RUN g++ -o app main.cpp


FROM alpine:latest
WORKDIR /app
COPY --from=build /app/app .
CMD ["./app"]