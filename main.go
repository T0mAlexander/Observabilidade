package main

import (
	"fmt"
	"go-store/routes"
	"net/http"
)

var serverPort = ":3333"

func main() {
	routes.Routes()
	fmt.Printf("Aplicação disponível no endereço http://localhost%s\n", serverPort)
	http.ListenAndServe(serverPort, nil)
}
