plugins {
  id("org.springframework.boot") version "3.2.2"
  id("io.spring.dependency-management") version "1.1.4"
  application
}

repositories {
  mavenCentral()
}

application {
  mainClass.set("com.example.app.AppStart")
}

dependencies {
  // Spring Boot
  implementation("org.springframework.boot:spring-boot-starter-web:3.2.2")
  implementation("org.springframework.boot:spring-boot-devtools:3.2.2")
  implementation("org.springframework.boot:spring-boot-starter-data-jpa:3.2.2")
  implementation("org.springframework.boot:spring-boot-starter-cache:3.2.2")
  implementation("org.springdoc:springdoc-openapi-webmvc-core:1.7.0")
  implementation("org.springdoc:springdoc-openapi-ui:1.7.0")

  // PostgreSQL Driver
  implementation("org.postgresql:postgresql:42.7.1")
  implementation("jakarta.persistence:jakarta.persistence-api:3.1.0")

  // HTTP Components
  implementation("org.apache.httpcomponents.client5:httpclient5:5.3.1")
  implementation("org.apache.httpcomponents.core5:httpcore5:5.2.4")
  implementation("org.apache.httpcomponents:fluent-hc:4.5.14")
  compileOnly("jakarta.servlet:jakarta.servlet-api:6.1.0-M1")

  // OpenTelemetry
  implementation("io.opentelemetry.javaagent:opentelemetry-javaagent:2.1.0")
  implementation("io.opentelemetry:opentelemetry-api:1.35.0")
  implementation("io.opentelemetry:opentelemetry-sdk-extension-autoconfigure:1.35.0")
  implementation("io.opentelemetry:opentelemetry-exporter-logging:1.35.0")
  implementation("io.opentelemetry:opentelemetry-exporter-otlp:1.31.0")
  implementation("io.opentelemetry:opentelemetry-exporter-common:1.31.0")
  implementation("io.micrometer:micrometer-tracing-bridge-otel:1.1.6")

  // Prometheus
  implementation("io.micrometer:micrometer-registry-prometheus:1.11.3")
  implementation("io.prometheus:prometheus-metrics-core:1.1.0")
  implementation("io.prometheus:prometheus-metrics-instrumentation-jvm:1.1.0")
  implementation("io.prometheus:prometheus-metrics-exporter-httpserver:1.1.0")
  implementation("io.prometheus:simpleclient_servlet:0.16.0")
  implementation("io.prometheus:simpleclient_httpserver:0.16.0")
}
