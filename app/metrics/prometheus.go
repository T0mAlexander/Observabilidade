package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	CpuTemperature prometheus.Gauge
	DiskFailures   *prometheus.CounterVec
	Registerer     prometheus.Registerer
}

var Registry = prometheus.NewRegistry()

func Prometheus() {
	// Registro global do Prometheus
	prom := NewMetrics(Registry)

	// Definindo valores para as métricas criadas
	prom.CpuTemperature.Set(65.3)
	prom.DiskFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	prometheus := &Metrics{
		CpuTemperature: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_temperature_celsius",
			Help: "Temperatura atual da CPU",
		}),
		DiskFailures: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "hd_errors_total",
				Help: "Total de erros de armazenamento",
			},
			[]string{"device"},
		),
	}
	reg.MustRegister(prometheus.CpuTemperature)
	reg.MustRegister(prometheus.DiskFailures)

	return prometheus
}
