FROM openjdk:alpine AS build
WORKDIR /app
COPY . .
RUN javac Main.java
CMD ["java", "-Xms512m", "-Xmx1024m", "-XX:LargePageSizeInBytes=1g", "-XX:ThreadStackSize=4m", "-XX:-UseContainerSupport", "Main"]