FROM gradle:8.6-alpine AS build
WORKDIR /app
COPY /app .
RUN gradle clean shadowJar

FROM openjdk:21-jdk

ENV TZ=America/Sao_Paulo

COPY --from=build /app/build/libs/app-all.jar /app-all.jar
COPY ./tools/otel-agent-2.2.0.jar /otel-agent.jar

CMD [ "java", "-javaagent:otel-agent.jar", "-jar", "app-all.jar" ]
