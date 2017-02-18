package handlers

import (
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server check. I'm alive!"))
}
