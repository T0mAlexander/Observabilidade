package com.example.routes;

import java.io.IOException;
import java.io.OutputStream;

import org.apache.http.client.fluent.Request;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;

public class Chain implements HttpHandler {
  Logger logger = LoggerFactory.getLogger(Chain.class);

  @Override
  public void handle(HttpExchange exchange) throws IOException {
    String HOST_ONE = System.getenv().getOrDefault("HOST_ONE", "localhost");
    String HOST_TWO = System.getenv().getOrDefault("HOST_ONE", "localhost");

    logger.debug("Comunicando com outras aplicações ...");
    Request.Get(String.format("http://%s:8080/io_task", HOST_ONE)).execute().returnContent();
    Request.Get(String.format("http://%s:8080/cpu_task", HOST_TWO)).execute().returnContent();

    logger.debug("Conexão encerrada");
    String bodyContent = "chain";
    exchange.getResponseHeaders().set("Content-Type", "text/plain");
    exchange.sendResponseHeaders(200, bodyContent.getBytes().length);
    try (OutputStream os = exchange.getResponseBody()) {
      os.write(bodyContent.getBytes());
    }
  }
}
