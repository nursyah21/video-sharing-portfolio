FROM amazoncorretto:17-alpine-jdk AS base
WORKDIR /app

COPY pom.xml mvnw ./
COPY .mvn .mvn
COPY src src

RUN ./mvnw clean install

FROM base

EXPOSE 5001
CMD [ "java", "-jar", "target/auth-0.0.1-SNAPSHOT.jar" ]