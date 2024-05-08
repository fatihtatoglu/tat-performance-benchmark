FROM buildo/scala-sbt-alpine:latest
WORKDIR /app
COPY . .
RUN scalac Main.scala
CMD ["scala", "Main"]
