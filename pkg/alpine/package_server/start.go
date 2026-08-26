package package_server

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func (s *Server) Start() {
	http.HandleFunc("/apk/", s.upload)
	go func() {
		errors.PanicOnError(http.ListenAndServe(":8080", nil))
	}()
	errors.PanicOnError(
		http.ListenAndServe(
			":8081",
			http.FileServer(http.Dir(constant.PackageRoot)),
		),
	)
}
