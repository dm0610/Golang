package controllers

import "net/http"

type userController struct{}

func (us userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}
