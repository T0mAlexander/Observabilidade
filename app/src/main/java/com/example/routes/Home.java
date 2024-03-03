package com.example.routes;

import java.io.IOException;
import java.io.OutputStream;
import java.util.List;
import java.util.Map;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.sun.net.httpserver.Headers;
import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;

public class Home implements HttpHandler {
  Logger logger = LoggerFactory.getLogger(Home.class);

  @Override
  public void handle(HttpExchange exchange) throws IOException {
    String name = "Java";
    Headers headers = exchange.getRequestHeaders();

    // Formatador e apresentador dos cabeçalhos da requisição
    StringBuilder headersInfo = new StringBuilder();
    headersInfo.append("Headers { ");
    for (Map.Entry<String, List<String>> entry : headers.entrySet()) {
      String key = entry.getKey();
      List<String> values = entry.getValue();
      headersInfo.append(key).append(": ").append(String.join(", ", values)).append("; ");
    }
    headersInfo.append("}");

    logger.error(headersInfo.toString());
    logger.error(String.format("Olá %s", name));
    logger.trace("Log de rastreamento");
    logger.debug("Log de depuração");
    logger.info("Log de informação");
    logger.warn("Log de aviso!");
    logger.error("Log de erro");

    String bodyContent = String.format("Olá %s!", name);
    exchange.getResponseHeaders().set("Content-Type", "text/plain");
    exchange.sendResponseHeaders(200, bodyContent.getBytes().length);
    try (OutputStream os = exchange.getResponseBody()) {
      os.write(bodyContent.getBytes());
    }
  }
}
