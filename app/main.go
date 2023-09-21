package main

import (
	"fmt"
	"go-store/metrics"
	"go-store/routes"
	"log"
	"net/http"
)

const serverPort = ":3333"

func main() {
	metrics.Data()
	routes.Routes()
	fmt.Printf("Aplicação disponível no endereço http://localhost%s\n", serverPort)
	log.Fatalln(http.ListenAndServe(serverPort, nil))
}
