package routes

import (
	"app/controllers"
	"app/metrics"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create", controllers.Product)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)

	http.Handle("/metrics", promhttp.HandlerFor(
		metrics.Prom,
		promhttp.HandlerOpts{
			Registry:           metrics.Prom,
			EnableOpenMetrics:  true,
			ErrorLog:           log.Default(),
			DisableCompression: true,
			ErrorHandling:      promhttp.HTTPErrorOnError,
		},
	))
}
