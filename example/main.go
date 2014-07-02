package main

import (
	"github.com/mgutz/gorazor/example/models"
	"github.com/mgutz/gorazor/example/views"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{Name: "Foo"}
	views.Index(user).WriteTo(w)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
