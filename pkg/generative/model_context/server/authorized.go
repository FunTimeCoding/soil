package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
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
	address := web.ClientAddress(r)

	if s.openAuthentication && t != "" && s.validateOpenToken(t) {
		console.Format("Authorized OIDC address:%s\n", address)

		return true
	}

	console.Format("Unauthorized token:%s address:%s\n", t, address)

	return false
}
