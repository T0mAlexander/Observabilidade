package com.example.routes;

import java.io.IOException;
import java.util.Arrays;
import java.util.List;
import java.util.Random;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;

public class RandomStatus implements HttpHandler {
  Logger logger = LoggerFactory.getLogger(RandomStatus.class);

  @Override
  public void handle(HttpExchange exchange) throws IOException {
    List<Integer> givenList = Arrays.asList(200, 200, 300, 400, 500);
    Random rand = new Random();
    int randomStatus = givenList.get(rand.nextInt(givenList.size()));

    logger.info("Status HTTP aleatório: " + randomStatus);
    String bodyContent = "Status HTTP: " + randomStatus;
    exchange.getResponseHeaders().set("Content-Type", "text/plain");
    exchange.sendResponseHeaders(randomStatus, bodyContent.getBytes().length);
    exchange.getResponseBody().write(bodyContent.getBytes());
    exchange.getResponseBody().close();
  }
}
