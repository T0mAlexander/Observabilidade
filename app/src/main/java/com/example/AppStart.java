package com.example;

import static java.lang.Thread.sleep;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.util.Random;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.example.routes.Chain;
import com.example.routes.CpuTask;
import com.example.routes.ErrorTest;
import com.example.routes.Home;
import com.example.routes.IOTask;
import com.example.routes.RandomStatus;
import com.sun.net.httpserver.HttpServer;

import io.prometheus.metrics.core.metrics.Counter;
import io.prometheus.metrics.core.metrics.Gauge;
import io.prometheus.metrics.core.metrics.Histogram;
import io.prometheus.metrics.core.metrics.Info;
import io.prometheus.metrics.core.metrics.StateSet;
import io.prometheus.metrics.core.metrics.Summary;
import io.prometheus.metrics.exporter.httpserver.HTTPServer;
import io.prometheus.metrics.instrumentation.jvm.JvmMetrics;
import io.prometheus.metrics.model.snapshots.Unit;

public class AppStart {
  public static void main(String[] args) throws IOException, Exception {
    // Provedor de logs
    Logger logger = LoggerFactory.getLogger(AppStart.class);
    HttpServer application = HttpServer.create(new InetSocketAddress(8080), 0);

    // Rotas da aplicação
    application.createContext("/", new Home());
    application.createContext("/io_task", new IOTask());
    application.createContext("/cpu_task", new CpuTask());
    application.createContext("/random_status", new RandomStatus());
    application.createContext("/chain", new Chain());
    application.createContext("/error_test", new ErrorTest());

    // Prometheus
    HTTPServer prometheusServer = HTTPServer.builder()
      .port(9464)
      .buildAndStart();

    // Métricas da JVM
    // Estas métricas são geradas e expostas automaticamente no endpoint do Prometheus
    JvmMetrics.builder().register();

    // Referência geral de tipos de métrica do Prometheus: https://tinyurl.com/prometheus-metric-types

    // Contador exemplar
    // Usados frequentemente para representar dados que progridem de forma aritmética/cumulativa
    // Exemplos de uso: contador de eventos, erros, excessões e requisições HTTP
    Counter counter = Counter.builder()
      .name("application_uptime_seconds")                               // Nome da métrica
      .help("Tempo total de atividade da aplicação em segundos")        // Descrição da métrica
      // .labelNames("rótulo")                                               // Rótulo da métrica
      .unit(Unit.SECONDS)                                                    // Unidade de medida
      .register();                                                           // Envio da métrica

    // Medidor/Manômetro exemplar
    // Representa um valor numérico variável ao longo da progressão do tempo
    // Exemplos de uso: consumo de memória, capacidade de armazenamento e uso da CPU
    Gauge gauge = Gauge.builder()
    .name("server_temperature_celsius")
    .help("Temperatura em graus Celsius")
    .unit(Unit.CELSIUS)
    .labelNames("localidade")
    .register();

    // Atribuição estática exemplar de valores
    gauge.labelValues("interna").set(23.4);
    gauge.labelValues("externa").set(9.3);

    //=====================================================================================================

    // Histograma exemplar
    // São usados para representar distribuição de frequências de forma cumulativa, simétrica e assimétrica
    // Exemplos de uso: latência de banco de dados, tamanho e tempo de resposta de requisições HTTP
    Histogram histogram = Histogram.builder()
    .name("application_request_latency_seconds")
    .help("Latência da requisição em segundos")
    .unit(Unit.SECONDS)
    .labelNames("status_http")
    .register();

    Random random = new Random(0);
    for (int i = 0; i < 1000; i++) {
      histogram.labelValues("200").observe(random.nextGaussian());
    }

    //=====================================================================================================

    // Sumário exemplar
    // Usado em cálculos estatísticos sobre distribuição de valores em função do tempo
    // Exemplos de uso: Analise de latência de consulta de banco de dados e monitoramento de tarefas assíncronas 
    Summary summary = Summary.builder()
    .name("application_response_latency_seconds")
    .help("Latência de resposta em segundos")
    .unit(Unit.BYTES)
    .quantile(0.95)
    .quantile(0.99)
    .register();

    for (int i = 0; i < 1000; i++) {
      summary.observe(random.nextGaussian());
    }

    //=====================================================================================================

    // Info exemplar
    // Utilizado para armazenar informações gerais da aplicação
    // Caso de uso: Informações sobre versão e metadados, identificação de componentes e dependências
    Info targetInfo = Info.builder()
      .name("target_info")
      .help("Recurso do OpenTelemetry")
      .labelNames("service.version")
      .register();

    targetInfo.setLabelValues("1.0.0");

    Info scopeInfo = Info.builder()
      .name("otel_scope_info")
      .labelNames("otel.scope.name", "otel.scope.version", "library_mascot")
      .register();

    scopeInfo.setLabelValues("minha.dependência", "100.3", "bear");

    Info info = Info.builder()
      .name("java_runtime_info")
      .help("Informações do JRE")
      .labelNames("version", "vendor", "runtime")
      .register();

    String version = System.getProperty("java.runtime.version", "desconhecido");
    String vendor = System.getProperty("java.vm.vendor", "desconhecido");
    String runtime = System.getProperty("java.runtime.name", "desconhecido");

    info.setLabelValues(version, vendor, runtime);

    //=====================================================================================================

    // Conjunto de estados exemplar
    // Usado para registro de conjunto de dados (carece de dados)
    // Possíveis exemplos de uso: controle de features do sistema, gestão de configs dinâmicas e representação de conjuntos enumerados/predefinidos
    StateSet stateSet = StateSet.builder()
    .name("feature_flags")
    .labelNames("env")
    .states("feature1", "feature2")
    .register();

    stateSet.labelValues("dev").setFalse("feature1");
    stateSet.labelValues("dev").setTrue("feature2");

    // Exportador do Open Telemetry
    // OpenTelemetryExporter.builder()
    //   .intervalSeconds(5)
    //   .registry(prometheusRegistry)
    //   .buildAndStart();

    // Inicialização da aplicação Java
    application.start();
    logger.info("Servidor iniciado no endereço http://localhost:" + application.getAddress().getPort());
    logger.info("Prometheus disponível no endereço http://localhost:" + prometheusServer.getPort());

    while (true) {
      sleep(1000);
      counter.inc();
    }
  }
}
