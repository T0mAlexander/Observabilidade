// Fornecedor automatizado de logs, métricas e traces do OpenTelemetry

package com.example.app;

import io.prometheus.client.exemplars.tracer.otel_agent.OpenTelemetryAgentSpanContextSupplier;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class TracingConfig {
  @Bean
  public OpenTelemetryAgentSpanContextSupplier OTelAgentSupplier() {
    return new OpenTelemetryAgentSpanContextSupplier();
  }
}
