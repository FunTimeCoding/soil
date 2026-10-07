package session

import (
	"github.com/funtimecoding/soil/pkg/provision/types/salt_login_request"
	"github.com/funtimecoding/soil/pkg/web/requester"
)

func New(
	login *requester.Requester,
	user string,
	password string,
	eauth string,
) *Session {
	return &Session{
		login: login,
		request: salt_login_request.Request{
			Username: user,
			Password: password,
			EAuth:    eauth,
		},
	}
}
