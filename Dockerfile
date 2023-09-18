FROM golang:1.21.1

WORKDIR /app
COPY /app .
RUN go build
RUN rm go-store

EXPOSE 3333
CMD [ "go", "run", "main.go" ]