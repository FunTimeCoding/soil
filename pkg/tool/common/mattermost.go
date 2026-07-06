package common

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/user_map"
)

func Mattermost(o ...mattermost.Option) *mattermost.Client {
	return mattermost.NewEnvironment(
		append(
			o,
			mattermost.WithUserMap(
				user_map.New(
					func(s string) string {
						return s
					},
					[]string{},
				),
			),
		)...,
	)
}
