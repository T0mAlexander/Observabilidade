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

tasks.withType<Jar> {
  manifest {
    attributes["Main-Class"] = "com.example.AppStart"
  }
}

dependencies {
  // OpenTelemetry
  implementation("io.opentelemetry:opentelemetry-exporter-otlp:1.36.0")

  // Prometheus
  implementation("io.prometheus:prometheus-metrics-core:1.1.0")
  implementation("io.prometheus:prometheus-metrics-exporter-httpserver:1.1.0")
  implementation("io.prometheus:prometheus-metrics-instrumentation-jvm:1.1.0")
  implementation("io.prometheus:prometheus-metrics-exporter-opentelemetry:1.1.0")

  // Logs
  implementation("org.slf4j:slf4j-api:2.0.12")
  implementation("ch.qos.logback:logback-core:1.5.0")
  implementation("ch.qos.logback:logback-classic:1.5.0")

  // Componentes HTTP
  implementation("org.apache.httpcomponents:fluent-hc:4.5.13")
}
