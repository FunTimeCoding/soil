package client

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (c *Client) SignOut(
	w http.ResponseWriter,
	r *http.Request,
) {
	http.SetCookie(
		w,
		&http.Cookie{
			Name:   constant.AuthorizationSubjectCookie,
			MaxAge: -1,
			Path:   constant.RootPath,
		},
	)
	locator := c.endSessionLocator()

	if locator == "" {
		http.Redirect(w, r, constant.RootPath, http.StatusFound)

		return
	}

	http.Redirect(w, r, locator, http.StatusFound)
}
