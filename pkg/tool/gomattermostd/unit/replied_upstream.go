package unit

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
)

func repliedUpstream(m *http.ServeMux) {
	root := &model.Post{Id: "alfa", ChannelId: "bravo", ReplyCount: 2}
	m.HandleFunc(
		"/api/v4/posts/alfa",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(w, root)
		},
	)
	m.HandleFunc(
		"/api/v4/posts/alfa/thread",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(
				w,
				&model.PostList{
					Order: []string{"delta", "charlie", "alfa"},
					Posts: map[string]*model.Post{
						"alfa": root,
						"charlie": {
							Id:       "charlie",
							RootId:   "alfa",
							UserId:   "foxtrot",
							Message:  "earlier reply",
							CreateAt: 1000,
						},
						"delta": {
							Id:       "delta",
							RootId:   "alfa",
							UserId:   "foxtrot",
							Message:  "latest reply",
							CreateAt: 2000,
						},
					},
				},
			)
		},
	)
	m.HandleFunc(
		"/api/v4/users/foxtrot",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			web.Encode(w, &model.User{Id: "foxtrot", Username: "foxtrot"})
		},
	)
}
