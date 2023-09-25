package controllers

import (
	"app/metrics"
	"app/models"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))
var Metrics http.Handler = MetricsSource()

func Index(write http.ResponseWriter, request *http.Request) {
	allProducts := models.FindAllProducts()

	error := temp.ExecuteTemplate(write, "Index", allProducts)

	if error != nil {
		log.Fatalln(error)
	}
}

func Product(write http.ResponseWriter, request *http.Request) {
	temp.ExecuteTemplate(write, "Create", nil)
}

func Insert(write http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		name := request.FormValue("name")
		description := request.FormValue("description")
		price := request.FormValue("price")
		amount := request.FormValue("amount")

		convertedPrice, error := strconv.ParseFloat(price, 64)
		if error != nil {
			log.Println("Error converting price:", error)
		}

		convertedAmount, error := strconv.Atoi(amount)
		if error != nil {
			log.Println("Error converting amount:", error)
		}

		models.CreateProduct(name, description, convertedPrice, convertedAmount)
	}
	http.Redirect(write, request, "/", http.StatusMovedPermanently)
}

func Delete(write http.ResponseWriter, request *http.Request) {
	productID := request.URL.Query().Get("id")
	models.DeleteProduct(productID)

	http.Redirect(write, request, "/", http.StatusMovedPermanently)
}

func Edit(write http.ResponseWriter, request *http.Request) {
	productID := request.URL.Query().Get("id")
	product := models.EditProduct(productID)

	temp.ExecuteTemplate(write, "Edit", product)
}

func Update(write http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		id := request.FormValue("id")
		name := request.FormValue("name")
		price := request.FormValue("price")
		description := request.FormValue("description")
		amount := request.FormValue("amount")

		convertedId, error := strconv.Atoi(id)
		if error != nil {
			log.Println("Error converting product ID to an integer:", error)
		}

		convertedPrice, error := strconv.ParseFloat(price, 64)
		if error != nil {
			log.Println("Error converting product price to a float64:", error)
		}

		convertedAmount, error := strconv.Atoi(amount)
		if error != nil {
			log.Println("Error converting product amount to an integer:", error)
		}

		models.UpdateProduct(convertedId, name, convertedPrice, description, convertedAmount)
	}

	http.Redirect(write, request, "/", http.StatusMovedPermanently)
}

func MetricsSource() http.Handler {
	return promhttp.HandlerFor(
		metrics.Prom,
		promhttp.HandlerOpts{
			Registry:          metrics.Prom,
			EnableOpenMetrics: true,
			ErrorLog: log.Default(),
			DisableCompression: true,
			ErrorHandling: promhttp.HTTPErrorOnError,
		},
	)
}
