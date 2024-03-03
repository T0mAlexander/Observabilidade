package com.example.routes;

import java.io.IOException;
import java.io.OutputStream;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;

public class IOTask implements HttpHandler {
  Logger logger = LoggerFactory.getLogger(IOTask.class);

  @Override
  public void handle(HttpExchange exchange) throws IOException {
    try {
      Thread.sleep(1000);
      logger.info("Tarefa da CPU");
    } catch (InterruptedException e) {
      Thread.currentThread().interrupt();
    }

    String bodyContent = "io_task";
    exchange.getResponseHeaders().set("Content-Type", "text/plain");
    exchange.sendResponseHeaders(200, bodyContent.getBytes().length);
    try (OutputStream os = exchange.getResponseBody()) {
      os.write(bodyContent.getBytes());
    }
  }
}
