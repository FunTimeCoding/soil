package server

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"net/http"
)

func New(
	i *identity.Tool,
	address string,
	setup func(*http.ServeMux),
) *Server {
	return &Server{
		Mux:      http.NewServeMux(),
		Setup:    setup,
		Address:  address,
		identity: i,
	}
}
