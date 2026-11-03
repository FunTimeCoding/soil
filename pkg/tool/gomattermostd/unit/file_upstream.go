package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
)

func fileUpstream(m *http.ServeMux) {
	m.HandleFunc(
		"/api/v4/files/alfa/info",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(w, &model.FileInfo{Id: "alfa", Name: "alfa.txt"})
		},
	)
	m.HandleFunc(
		"/api/v4/files/alfa",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			_, e := w.Write([]byte("alfa"))
			errors.PanicOnError(e)
		},
	)
}
