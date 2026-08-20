package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"time"
)

func (c *Client) SubjectCookie(subject string) *http.Cookie {
	encrypted, e := c.encrypt([]byte(subject))
	errors.PanicOnError(e)

	return &http.Cookie{
		Name:     constant.AuthorizationSubjectCookie,
		Value:    encrypted,
		HttpOnly: true,
		Secure:   true,
		Path:     constant.RootPath,
		Expires:  time.Now().Add(24 * time.Hour),
	}
}
