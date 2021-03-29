package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - almacena pointer para servicio dado
type Handler struct {
	Router *mux.Router
}

// NewHnadler - retorna un pointer a Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - sets ups all routes: router.
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Esto vivo!")
	})
}
