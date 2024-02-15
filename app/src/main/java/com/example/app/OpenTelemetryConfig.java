// Configuração do OpenTelemetry
// Referência técnica: https://github.com/open-telemetry/opentelemetry-java-examples/tree/main/autoconfigure

package com.example.app;

import io.opentelemetry.sdk.autoconfigure.AutoConfiguredOpenTelemetrySdk;

public final class OpenTelemetryConfig {
  public static void main(String[] args) throws InterruptedException {
    AutoConfiguredOpenTelemetrySdk.initialize().getOpenTelemetrySdk();
  }
}
