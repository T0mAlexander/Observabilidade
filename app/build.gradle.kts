plugins {
  java
  application
  id("com.github.johnrengelman.shadow") version "8.1.1"
}

repositories {
  mavenCentral()
}

application {
  mainClass.set("com.example.AppStart")
}

tasks.withType<JavaExec> {
  jvmArgs(
    "-javaagent:../tools/otel-agent.jar",
    "-Dotel.service.name=app-java",
    "-Dotel.logs.exporter=otlp",
    "-Dotel.metrics.exporter=otlp",
    "-Dotel.traces.exporter=otlp",
    "-Dotel.traces.sampler=always_on"
    // "-Dotel.exporter.otlp.metrics.endpoint=http://localhost:9090/api/v1/write",
    // "-Dotel.exporter.otlp.logs.endpoint=http://localhost:3100/api/prom/push",
    // "-Dotel.exporter.otlp.traces.endpoint=http://localhost:4317",
    // "-Dotel.exporter.otlp.endpoint=http://localhost:4317",
  )
}

tasks.withType<Jar> {
  manifest {
    attributes["Main-Class"] = "com.example.AppStart"
  }
}

dependencies {
  // Componentes HTTP
  implementation("org.apache.httpcomponents:fluent-hc:4.5.13")

  // Prometheus
  implementation("io.prometheus:prometheus-metrics-core:1.1.0")
  implementation("io.prometheus:prometheus-metrics-exporter-httpserver:1.1.0")
  implementation("io.prometheus:prometheus-metrics-instrumentation-jvm:1.1.0")
  implementation("io.prometheus:prometheus-metrics-exporter-opentelemetry:1.1.0")
  implementation("io.micrometer:micrometer-registry-prometheus:1.12.4")

  // Logs
  implementation("org.slf4j:slf4j-api:2.0.12")
  implementation("ch.qos.logback:logback-core:1.5.0")
  implementation("ch.qos.logback:logback-classic:1.5.0")
}
