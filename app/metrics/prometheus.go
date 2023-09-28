package metrics

import (
	"app/controllers"

	"github.com/prometheus/client_golang/prometheus"
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
		metrics.CpuTemperature,
		controllers.RequestsTotal,
	)

	return metrics
}
