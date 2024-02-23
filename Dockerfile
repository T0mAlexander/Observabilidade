FROM curlimages/curl:8.6.0 AS download
ARG OTEL_VERSION="v2.1.0"
RUN curl -fsL "https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/download/$OTEL_VERSION/opentelemetry-javaagent.jar" \
    -o "$HOME/otel-agent.jar"

FROM gradle:8.6.0-jdk21-alpine AS build
WORKDIR /app
COPY /app .
RUN gradle build

FROM openjdk:21-jdk
COPY --from=build /app/build/libs/app.jar /app.jar
COPY --from=download /home/curl_user/otel-agent.jar /otel-agent.jar

ENTRYPOINT [ "java", "-javaagent:/otel-agent.jar", "-jar", "/app.jar" ]