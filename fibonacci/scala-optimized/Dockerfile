FROM buildo/scala-sbt-alpine:latest
WORKDIR /app
COPY . .
RUN scalac -g:none -optimize Main.scala
CMD ["scala", "Main"]
