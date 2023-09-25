package metrics

import (
	"app/routes/middleware"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	CpuTemperature prometheus.Gauge
	RequestsTotal  prometheus.CounterVec
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
		metrics.CpuTemperature,
		middleware.RequestsTotal,
	)

	return metrics
}
