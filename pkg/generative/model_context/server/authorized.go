package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"strings"
)

func (s *Server) authorized(r *http.Request) bool {
	t := strings.TrimPrefix(
		r.Header.Get(webConstant.Authorization),
		fmt.Sprintf("%s ", webConstant.Bearer),
	)

	if t == "" {
		t = r.URL.Query().Get(generative.ModelContextTokenParameter)
	}

	address := web.ClientAddress(r)

	if s.tokenAuthentication && s.token != "" && t == s.token {
		console.Format("Authorized token:%s address:%s\n", t, address)

		return true
	}

	if s.openAuthentication && t != "" && s.validateOpenToken(t) {
		console.Format("Authorized OIDC address:%s\n", address)

		return true
	}

	console.Format("Unauthorized token:%s address:%s\n", t, address)

	return false
}
