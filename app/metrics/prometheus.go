package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Metrics struct {
	CpuTemperature prometheus.Gauge
}

var Prom = prometheus.NewRegistry()

func Prometheus() *Metrics {
	metrics := &Metrics{
		CpuTemperature: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "cpu_temperature_celsius",
				Help: "Temperatura atual da CPU",
			}),
	}

	Prom.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(
			collectors.ProcessCollectorOpts{
				ReportErrors: true,
			},
		),

		metrics.CpuTemperature,
	)

	return metrics
}
