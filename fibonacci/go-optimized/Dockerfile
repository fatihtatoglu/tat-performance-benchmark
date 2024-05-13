FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go build -a -gcflags=all="-l -B" -ldflags="-w -s" -o main .


FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
CMD [ "./main" ]