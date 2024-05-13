FROM rust:alpine AS build
WORKDIR /app
COPY . .
RUN cargo build

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/target/debug/app .
CMD ["./app"]
