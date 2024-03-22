FROM gradle:8.6-alpine AS build

WORKDIR /app
COPY /app .
RUN gradle clean build

FROM jboss/wildfly:13.0.0.Final

ENV WILDFLY_DIR="/opt/jboss/wildfly"

COPY ./tools/opentelemetry-agent-1.33.1.jar /opentelemetry-agent-1.33.1.jar
COPY ./tools/jmx-agent-0.20.0.jar /jmx-agent-0.20.0.jar

RUN sh /opt/jboss/wildfly/bin/add-user.sh --user "admin" --password "admin" --silent
COPY --from=build /app/build/libs/app.war /opt/jboss/wildfly/standalone/deployments/app.war

EXPOSE 8080 9990 9991
