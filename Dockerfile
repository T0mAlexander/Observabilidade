FROM gradle:8.6-alpine AS build
WORKDIR /app
COPY /app .
RUN gradle clean shadowJar

# Ocelot suporta no máximo OpenJDK 20
FROM openjdk:20-jdk
COPY --from=build /app/build/libs/app-all.jar /app-all.jar
COPY ./tools/ocelot-agent-2.6.4.jar /ocelot-agent-2.6.4.jar

CMD [ "java", "-javaagent:ocelot-agent-2.6.4.jar", "-Dinspectit.config.file-based.path=/app/config", "-jar", "app-all.jar" ]
