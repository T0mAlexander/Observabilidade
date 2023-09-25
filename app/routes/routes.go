package routes

import (
	"app/controllers"
	"app/metrics"
	"app/routes/middleware"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create", controllers.Product)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)

	http.Handle(
		"/metrics", middleware.Prometheus(
			metrics.Prom, nil).Middleware(
			"/metrics", controllers.Metrics,
		))
}
