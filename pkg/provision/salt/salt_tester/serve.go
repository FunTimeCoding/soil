package salt_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"strings"
)

func (s *Server) serve(
	w http.ResponseWriter,
	r *http.Request,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	path := strings.TrimPrefix(r.URL.Path, webConstant.RootPath)

	if path == constant.SaltLoginPath {
		s.login(w, r)

		return
	}

	token := r.Header.Get(constant.SaltTokenHeader)
	s.seen = append(s.seen, token)

	if token != s.valid {
		w.WriteHeader(http.StatusUnauthorized)
		write(
			w,
			"<html><p>No permission -- see authorization schemes</p></html>",
		)

		return
	}

	if path == constant.SaltKeysPath {
		write(w, `{"return":{"minions":["alfa"],"minions_pre":[]}}`)

		return
	}

	if path == "" && len(s.answers) > 0 {
		write(w, fmt.Sprintf(`{"return":[%s]}`, notation.Marshal(s.answers)))

		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	write(w, "<html><p>An unexpected error occurred</p></html>")
}
