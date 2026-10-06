package session

import "github.com/funtimecoding/soil/pkg/web/requester"

func New(
	login *requester.Requester,
	user string,
	password string,
	eauth string,
) *Session {
	return &Session{
		login: login,
		request: loginRequest{
			Username: user,
			Password: password,
			EAuth:    eauth,
		},
	}
}
