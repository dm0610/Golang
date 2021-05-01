package prom_controller

import (
	"fmt"
	"net/http"
)

func RegisterControllers() {
	uc := metrics()
	http.Handle("/metrics", *uc)
}

func metrics() {
	fmt.Println("Hello World!")
}
