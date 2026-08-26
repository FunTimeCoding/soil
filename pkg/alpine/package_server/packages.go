package package_server

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (s *Server) packages(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

		return
	}

	listings, e := Indexes(constant.PackageRoot)

	if e != nil {
		httpFail(w, "read indexes fail", e)

		return
	}

	web.EncodeNotation(w, listings)
}
