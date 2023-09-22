package main

import (
	"app/metrics"
	"app/routes"
	"fmt"
	"log"
	"net/http"
)

const serverPort = ":3333"

func main() {
	metrics.Cpu()
	routes.Routes()
	fmt.Printf("Aplicação disponível no endereço http://localhost%s\n", serverPort)
	log.Fatalln(http.ListenAndServe(serverPort, nil))
}
