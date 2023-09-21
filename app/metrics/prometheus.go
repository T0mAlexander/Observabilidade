package metrics

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Metrics struct {
	CpuTemperature prometheus.Gauge
	Latency        prometheus.Histogram
	Registerer     prometheus.Registerer
}

var Prometheus = prometheus.NewRegistry()

func Data() {
	// Registro global de métricas do Prometheus
	prom := NewMetrics()

	// Rotina assíncrona para gerar temperatura fictícia
	go func() {
		// Temporizador de 3 segundos
		timer := time.NewTicker(3 * time.Second)

		for {
			// Gerando algarismo aleatório entre 40 até 75
			randomNumber := rand.Float64()*(75.0-40.0) + 40.0

			// Limitando em 1 casa decimal pós ponto
			generatedNumber := fmt.Sprintf("%.1f", randomNumber)

			// Conversão de string em float64
			temperature, error := strconv.ParseFloat(generatedNumber, 64)

			// Tratamento de erro
			if error != nil {
				fmt.Println("Erro ao converter numeral", error)
			}

			// Informando a métrica fictícia ao Prometheus
			prom.CpuTemperature.Set(temperature)

			// Resetando o temporizador
			<-timer.C
		}
	} ()
}

func NewMetrics() *Metrics {
	metrics := &Metrics{
		CpuTemperature: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "cpu_temperature_celsius",
				Help: "Temperatura atual da CPU",
		}),

		Latency: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "api",
				Name:      "http_request_duration_seconds",
				Help:      "Histograma de duração de requisições HTTP em segundos",
				Buckets:   prometheus.ExponentialBuckets(0.1, 1.5, 5),
		}),
	}

	Prometheus.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(
			collectors.ProcessCollectorOpts{
				ReportErrors: true,
			},
		),

		metrics.CpuTemperature,
		metrics.Latency,
	)

	go func() {
		for {
			now := time.Now()
			metrics.Latency.(prometheus.ExemplarObserver).ObserveWithExemplar(
				time.Since(now).Seconds(),
				prometheus.Labels{
					"dummyID": fmt.Sprint(rand.Intn(100000)),
				},
			)

			time.Sleep(600 * time.Millisecond)
		}
	}()

	return metrics
}
