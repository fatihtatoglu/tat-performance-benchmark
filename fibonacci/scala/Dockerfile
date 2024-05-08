FROM hseeberger/scala-sbt:latest

WORKDIR /app

COPY . .

RUN scalac Main.scala

CMD ["scala", "Main"]
