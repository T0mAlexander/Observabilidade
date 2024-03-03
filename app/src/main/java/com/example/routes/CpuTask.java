package com.example.routes;

import java.io.IOException;
import java.io.OutputStream;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;

public class CpuTask implements HttpHandler {
  Logger logger = LoggerFactory.getLogger(CpuTask.class);

  @Override
  @SuppressWarnings("unused")
  public void handle(HttpExchange exchange) throws IOException {
    for (int x = 0; x < 100; x++) {
      int tmp = x * x * x;
    }

    logger.info("Tarefa da CPU");
    String bodyContent = "cpu_task";
    exchange.getResponseHeaders().set("Content-Type", "text/plain");
    exchange.sendResponseHeaders(200, bodyContent.getBytes().length);
    try (OutputStream os = exchange.getResponseBody()) {
      os.write(bodyContent.getBytes());
    }
  }
}
