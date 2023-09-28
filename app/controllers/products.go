package controllers

import (
	"app/models"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))
var RequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requisições HTTP",
	}, []string{"method", "route", "status"},
)

func Index(write http.ResponseWriter, request *http.Request) {
	allProducts := models.FindAllProducts()

	error := temp.ExecuteTemplate(write, "Index", allProducts)

	if error != nil {
		log.Fatalln(error)
		RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusBadRequest)).Inc()
	}

	RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusOK)).Inc()
}

func Product(write http.ResponseWriter, request *http.Request) {
	error := temp.ExecuteTemplate(write, "Create", nil)

	if error != nil {
		log.Fatalln(error)
		RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusUnauthorized)).Inc()
	}

	RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusCreated)).Inc()
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
			RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusUnprocessableEntity)).Inc()
		}

		convertedAmount, error := strconv.Atoi(amount)
		if error != nil {
			log.Println("Error converting amount:", error)
			RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusUnprocessableEntity)).Inc()
		}

		models.CreateProduct(name, description, convertedPrice, convertedAmount)
	}
	http.Redirect(write, request, "/", http.StatusMovedPermanently)
	RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusCreated)).Inc()
}

func Delete(write http.ResponseWriter, request *http.Request) {
	productID := request.URL.Query().Get("id")
	models.DeleteProduct(productID)

	http.Redirect(write, request, "/", http.StatusMovedPermanently)
	RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusNoContent)).Inc()
}

func Edit(write http.ResponseWriter, request *http.Request) {
	productID := request.URL.Query().Get("id")
	product := models.EditProduct(productID)

	RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusOK)).Inc()

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
			RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusUnprocessableEntity)).Inc()
		}

		convertedPrice, error := strconv.ParseFloat(price, 64)
		if error != nil {
			log.Println("Error converting product price to a float64:", error)
			RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusUnprocessableEntity)).Inc()
		}

		convertedAmount, error := strconv.Atoi(amount)
		if error != nil {
			log.Println("Error converting product amount to an integer:", error)
			RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusUnprocessableEntity)).Inc()
		}

		models.UpdateProduct(convertedId, name, convertedPrice, description, convertedAmount)
		RequestsTotal.WithLabelValues(request.Method, request.URL.Path, strconv.Itoa(http.StatusAccepted)).Inc()
	}

	http.Redirect(write, request, "/", http.StatusMovedPermanently)
}
