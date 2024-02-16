FROM curlimages/curl:8.6.0 AS download
ARG OTEL_AGENT_VERSION="1.32.1"
RUN curl -fsL "https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/download/v$OTEL_AGENT_VERSION/opentelemetry-javaagent.jar" \
    -o "$HOME/opentelemetry-agent.jar"

FROM maven:3.8.6-openjdk-8 AS build
ADD . /build
RUN cd /build && mvn package --quiet -DskipTests

FROM openjdk:8-jre-alpine
COPY --from=build /build/target/*.jar /app.jar
COPY --from=download /home/curl_user/opentelemetry-agent.jar /opentelemetry-agent.jar

ENTRYPOINT ["java", "-javaagent:/opentelemetry-agent.jar", "-jar", "/app.jar" ]
