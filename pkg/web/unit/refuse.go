package unit

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func refuse(http.HandlerFunc) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		q *http.Request,
	) {
		http.Redirect(w, q, constant.SignInPath, http.StatusFound)
	}
}
