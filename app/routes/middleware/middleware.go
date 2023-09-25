package middleware

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Middleware interface {
	Middleware(handler string, next http.Handler) http.HandlerFunc
}

type middleware struct {
	registry prometheus.Registerer
	buckets  []float64
}

var handler string
var registry prometheus.Registerer
var buckets []float64

var RequestsTotal = promauto.With(
	prometheus.WrapRegistererWith(
		prometheus.Labels{"handler": handler}, registry)).NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requisições HTTP",
	}, []string{"method", "code"},
)
var RequestDuration = promauto.With(registry).NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Latência das requisições HTTP",
		Buckets: buckets,
	},
	[]string{"method", "code"},
)
var RequestSize = promauto.With(registry).NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_size_bytes",
		Help: "Tamanho das requisições HTTP",
	}, []string{"method", "code"},
)
var ResponseSize = promauto.With(registry).NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_response_size_bytes",
		Help: "Tamanho das respostas HTTP",
	}, []string{"method", "code"},
)

func (*middleware) Middleware(handler string, next http.Handler) http.HandlerFunc {
	wrapper := promhttp.InstrumentHandlerCounter(
		RequestsTotal,
		promhttp.InstrumentHandlerDuration(
			RequestDuration,
			promhttp.InstrumentHandlerRequestSize(
				RequestSize,
				promhttp.InstrumentHandlerResponseSize(
					ResponseSize,
					http.HandlerFunc(
						func(writer http.ResponseWriter, response *http.Request) {
							next.ServeHTTP(writer, response)
						}),
				),
			),
		),
	)

	return wrapper.ServeHTTP
}

func Prometheus(registry prometheus.Registerer, buckets []float64) Middleware {
	if buckets == nil {
		buckets = prometheus.ExponentialBuckets(0.1, 1.5, 5)
	}

	return &middleware{
		buckets:  buckets,
		registry: registry,
	}
}
