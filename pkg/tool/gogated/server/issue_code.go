package server

import (
	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/token/jwt"
	"net/http"
	"time"
)

func (s *Server) issueCode(
	w http.ResponseWriter,
	r *http.Request,
	authorizeRequest fosite.AuthorizeRequester,
	subject string,
	authenticatedAt time.Time,
	requestedAt time.Time,
) {
	for _, scope := range authorizeRequest.GetRequestedScopes() {
		authorizeRequest.GrantScope(scope)
	}

	authorizeResponse, e := s.service.NewAuthorizeResponse(
		r,
		authorizeRequest,
		&openid.DefaultSession{
			Claims: &jwt.IDTokenClaims{
				Subject:     subject,
				AuthTime:    authenticatedAt,
				RequestedAt: requestedAt,
			},
			Subject: subject,
		},
	)

	if e != nil {
		s.service.WriteAuthorizeError(r, w, authorizeRequest, e)

		return
	}

	s.service.WriteAuthorizeResponse(r, w, authorizeRequest, authorizeResponse)
}
