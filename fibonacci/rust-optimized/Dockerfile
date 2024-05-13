FROM rust:alpine AS build
WORKDIR /app
COPY . .
RUN cargo build --release

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/target/release/app .
CMD ["./app"]
