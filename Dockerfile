FROM gradle:8.6.0-jdk21-alpine

WORKDIR /app
COPY ./app .

CMD [ "./gradlew", "bootRun" ]
