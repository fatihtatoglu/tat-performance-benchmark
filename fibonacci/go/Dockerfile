FROM golang:alpine

WORKDIR /usr/src/app

COPY . .

RUN go build -o main .

CMD [ "./main" ]