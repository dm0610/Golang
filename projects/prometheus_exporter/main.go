package main

import (
	"net/http"
)

func main() {

	prom_controller.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
