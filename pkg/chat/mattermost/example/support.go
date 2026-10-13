package example

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Support() {
	user1 := common.Mattermost(
		mattermost.WithToken(environment.Required("MATTERMOST_TOKEN_USER1")),
	)
	defer user1.Close()
	user2 := common.Mattermost(
		mattermost.WithToken(environment.Required("MATTERMOST_TOKEN_USER2")),
	)
	defer user2.Close()
	h := user1.DefaultChannel()
	t := user1.MustPostSimple(
		h,
		"I have a problem with my server, its not responding anymore",
	)
	user2.MustReply(h, t, "Can you try to restart it?")
}
