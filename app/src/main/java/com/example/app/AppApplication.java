package com.example.app;

import org.apache.http.client.fluent.Request;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cache.annotation.EnableCaching;
import org.springframework.http.HttpHeaders;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.Arrays;
import java.util.List;
import java.util.Random;

@SpringBootApplication
@RestController
@EnableCaching
public class AppApplication {
  @Autowired
  private PeanutsService service;

  Logger logger = LoggerFactory.getLogger(AppApplication.class);

  public static void main(String[] args) {
    SpringApplication.run(AppApplication.class, args);
  }

  @GetMapping("/")
  public String root(@RequestParam(value = "name", defaultValue = "Java") String name, @RequestHeader HttpHeaders headers) {
    logger.error(headers.toString());
    logger.error(String.format("Olá %s!", name));
    logger.debug("Log de depuração");
    logger.info("Log de informação");
    logger.warn("Log de aviso!");
    logger.error("Log de erro");
    return String.format("Olá %s!!", name);
  }

  @GetMapping("/io_task")
  public String io_task() throws InterruptedException {
    Thread.sleep(1000);
    logger.info("Tarefa I/O");
    return "io_task";
  }

  @SuppressWarnings("unused")
  @GetMapping("/cpu_task")
  public String cpu_task() {
    for (int i = 0; i < 100; i++) {
      int tmp = i * i * i;
    }
    logger.info("Tarefa da CPU");
    return "cpu_task";
  }

  @GetMapping("/random_sleep")
  public String random_sleep() throws InterruptedException {
    Thread.sleep((int) (Math.random() / 5 * 10000));
    logger.info("Ociosidade aleatória");
    return "random_sleep";
  }

  @GetMapping("/random_status")
  public String random_status(HttpServletResponse response) throws InterruptedException {
    List<Integer> givenList = Arrays.asList(200, 200, 300, 400, 500);
    Random rand = new Random();
    int randomElement = givenList.get(rand.nextInt(givenList.size()));
    response.setStatus(randomElement);
    logger.info("Status HTTP aleatório");
    return "random_status";
  }

  @GetMapping("/chain")
  public String chain() throws InterruptedException, IOException {
    String TARGET_ONE_HOST = System.getenv().getOrDefault("TARGET_ONE_HOST", "localhost");
    String TARGET_TWO_HOST = System.getenv().getOrDefault("TARGET_TWO_HOST", "localhost");
    logger.debug("Cadeia inicializando ...");
    Request.Get("http://localhost:8080/").execute().returnContent();
    Request.Get(String.format("http://%s:8080/io_task", TARGET_ONE_HOST)).execute().returnContent();
    Request.Get(String.format("http://%s:8080/cpu_task", TARGET_TWO_HOST)).execute().returnContent();
    logger.debug("Cadeia encerrada");
    return "chain";
  }

  @GetMapping("/error_test")
  public String error_test() throws Exception {
    throw new Exception("Erro de teste");
  }

  @GetMapping("/peanuts/{id}")
  public Peanuts getPeanutsById(@PathVariable Long id) {
    logger.info("Obtendo recursos");
    return service.getPeanutsById(id);
  }

  @PostMapping("/peanuts")
  public Peanuts savePeanuts(@RequestBody Peanuts peanuts) {
    logger.info("Criando recurso");
    return service.savePeanuts(peanuts);
  }
}
