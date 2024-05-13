FROM alpine:latest AS build
WORKDIR /app
COPY . .
RUN apk --update add g++
RUN g++ -o app -O3 main.cpp


FROM alpine:latest
WORKDIR /app
RUN apk --update add libstdc++
COPY --from=build /app/app .
CMD ["./app"]