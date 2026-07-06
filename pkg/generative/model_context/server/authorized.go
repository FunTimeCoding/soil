package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/model_context/constant"
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
		t = r.URL.Query().Get(constant.TokenParameter)
	}

	address := web.ClientAddress(r)

	if s.tokenAuthentication && s.token != "" && t == s.token {
		fmt.Printf(
			"Authorized token:%s address:%s\n",
			t,
			address,
		)

		return true
	}

	if s.openAuthentication && t != "" && s.validateOpenToken(t) {
		fmt.Printf("Authorized OIDC address:%s\n", address)

		return true
	}

	fmt.Printf(
		"Unauthorized token:%s address:%s\n",
		t,
		address,
	)

	return false
}
