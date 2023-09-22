package routes

import (
	"app/controllers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create", controllers.Product)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
	http.Handle("/metrics", controllers.Metrics)
}
