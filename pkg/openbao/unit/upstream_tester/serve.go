package upstream_tester

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/openbao/constant"
	"github.com/openbao/openbao/api/v2"
	"net/http"
	"strings"
)

func (u *Tester) serve(
	w http.ResponseWriter,
	r *http.Request,
) {
	u.lock.Lock()
	defer u.lock.Unlock()

	if !strings.HasPrefix(r.URL.Path, "/v1/sys/") {
		u.tokens = append(u.tokens, r.Header.Get(api.AuthHeaderName))
	}

	listing := r.Method == constant.ListMethod ||
		r.URL.Query().Get(constant.ListParameter) == "true"

	switch {
	case r.URL.Path == "/v1/sys/health":
		answer(
			w,
			http.StatusOK,
			`{"initialized":true,"sealed":false,"version":"2.7.1"}`,
		)
	case r.URL.Path == "/v1/auth/kubernetes/login":
		var body map[string]string
		notation.MustDecodeBytes(readBody(r), &body, false)

		if body["jwt"] != "alfa-jwt" || body["role"] != "alfa-role" {
			answer(w, http.StatusForbidden, `{"errors":["permission denied"]}`)

			return
		}

		u.logins++
		answer(
			w,
			http.StatusOK,
			`{"auth":{"client_token":"alfa-token","lease_duration":3600}}`,
		)
	case listing && r.URL.Path == "/v1/secret/metadata/alfa":
		answer(w, http.StatusOK, `{"data":{"keys":["charlie/","bravo"]}}`)
	case listing:
		answer(w, http.StatusNotFound, `{"errors":[]}`)
	case r.Method == http.MethodGet && r.URL.Path == "/v1/secret/data/alfa/bravo":
		answer(
			w,
			http.StatusOK,
			`{"data":{"data":{"DELTA":"echo","FOXTROT":"golf"},"metadata":{"version":3,"created_time":"2026-10-11T00:00:00Z","deletion_time":"","destroyed":false}}}`,
		)
	case r.Method == http.MethodGet:
		answer(w, http.StatusNotFound, `{"errors":[]}`)
	default:
		var body map[string]any
		notation.MustDecodeBytes(readBody(r), &body, false)
		u.written = body
		answer(
			w,
			http.StatusOK,
			`{"data":{"version":1,"created_time":"2026-10-11T00:00:00Z","deletion_time":"","destroyed":false}}`,
		)
	}
}
