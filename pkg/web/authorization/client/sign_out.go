package client

import (
	"github.com/funtimecoding/soil/pkg/web/authorization/constant"
	"net/http"
)

func (c *Client) SignOut(
	w http.ResponseWriter,
	r *http.Request,
) {
	http.SetCookie(
		w,
		&http.Cookie{
			Name:   constant.SubjectCookie,
			MaxAge: -1,
			Path:   "/",
		},
	)
	http.Redirect(w, r, "/", http.StatusFound)
}
