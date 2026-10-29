package unit

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
)

func unresolvedUpstream(m *http.ServeMux) {
	serve := func(
		path string,
		v any,
	) {
		m.HandleFunc(
			path,
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				web.Encode(w, v)
			},
		)
	}
	resolved := &model.Post{
		Id:           "bravo",
		UserId:       "delta",
		Message:      "resolved question",
		CreateAt:     2000,
		HasReactions: true,
	}
	open := &model.Post{
		Id:         "kilo",
		UserId:     "delta",
		Message:    "open question",
		CreateAt:   3000,
		ReplyCount: 1,
	}
	reply := &model.Post{
		Id:       "lima",
		UserId:   "delta",
		RootId:   "kilo",
		Message:  "still looking",
		CreateAt: 4000,
	}
	serve(
		"/api/v4/teams/tango/channels/name/alfa",
		&model.Channel{Id: "alfa", Name: "alfa"},
	)
	serve(
		"/api/v4/channels/alfa/posts",
		&model.PostList{
			Order: []string{"lima", "kilo", "bravo"},
			Posts: map[string]*model.Post{
				"bravo": resolved,
				"kilo":  open,
				"lima":  reply,
			},
		},
	)
	serve(
		"/api/v4/posts/bravo/reactions",
		[]*model.Reaction{
			{PostId: "bravo", EmojiName: constant.MattermostDoneReaction},
		},
	)
	serve(
		"/api/v4/posts/kilo/thread",
		&model.PostList{
			Order: []string{"lima", "kilo"},
			Posts: map[string]*model.Post{"kilo": open, "lima": reply},
		},
	)
	serve("/api/v4/users/delta", &model.User{Id: "delta", Username: "echo"})
}
