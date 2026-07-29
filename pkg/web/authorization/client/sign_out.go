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
			Path:   constant.LocationRoot,
		},
	)
	http.Redirect(w, r, constant.LocationRoot, http.StatusFound)
}
