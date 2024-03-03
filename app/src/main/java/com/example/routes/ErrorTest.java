package com.example.routes;

import java.io.IOException;
import java.io.OutputStream;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;

public class ErrorTest implements HttpHandler {
  Logger logger = LoggerFactory.getLogger(ErrorTest.class);

  @Override
  public void handle(HttpExchange exchange) throws IOException {
    try {
      throw new Exception("Erro de teste");
    } catch (Exception error) {
      logger.error("Erro de teste", error);
      
      String errorMessage = error.getMessage();
      exchange.sendResponseHeaders(500, errorMessage.length());
      try (OutputStream os = exchange.getResponseBody()) {
        os.write(errorMessage.getBytes());
      }
    }
  }
}
