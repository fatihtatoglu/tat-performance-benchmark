FROM openjdk:alpine AS build
WORKDIR /app
COPY . .
RUN javac Main.java
CMD ["java", "Main"]