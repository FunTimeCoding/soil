package model_context

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
)

func upstream(m *http.ServeMux) {
	m.HandleFunc(
		"/api/v4/posts/alfa",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(w, &model.Post{Id: "alfa", ChannelId: "bravo"})
		},
	)
	m.HandleFunc(
		"/api/v4/posts/reply",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(
				w,
				&model.Post{
					Id:        "reply",
					RootId:    "alfa",
					ChannelId: "bravo",
				},
			)
		},
	)
}
