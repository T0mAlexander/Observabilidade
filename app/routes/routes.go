package routes

import (
	"go-store/controllers"
	"go-store/metrics"
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
	http.Handle("/metrics", promhttp.HandlerFor(metrics.Prometheus, promhttp.HandlerOpts{Registry: metrics.Prometheus}))
}
