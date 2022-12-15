package main

import (
	"fmt"
	"net/http"
)

func getUserProfile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "text/plain")

	//búsqueda en la base de datos

	//devuelve el usuario buscado, error si no lo encuentra
}
func main() {
	http.HandleFunc("/", getUserProfile)
	fmt.Println("Starting web server at http://localhost:8090")
	http.ListenAndServe(":8090", nil)
}
