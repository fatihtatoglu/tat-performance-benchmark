FROM rust:alpine AS build

WORKDIR /usr/src/app

COPY . .

RUN cargo build --release

FROM alpine:latest
WORKDIR /usr/src/app
COPY --from=build /usr/src/app/target/release/app .
CMD ["./app"]
