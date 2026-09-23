package worker

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
)

func upstream(m *http.ServeMux) {
	m.HandleFunc(
		"/api/v4/users/me",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(w, &model.User{Id: "self", Username: "assistant"})
		},
	)
	m.HandleFunc(
		"/api/v4/users/foxtrot",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(w, &model.User{Id: "foxtrot", Username: "Foxtrot"})
		},
	)
	m.HandleFunc(
		"/api/v4/posts/alfa",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(w, &model.Post{Id: "alfa", ChannelId: "bravo"})
		},
	)
}
