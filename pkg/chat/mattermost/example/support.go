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
	user2 := common.Mattermost(
		mattermost.WithToken(environment.Required("MATTERMOST_TOKEN_USER2")),
	)
	h := user1.DefaultChannel()
	t := user1.MustPostSimple(
		h,
		"I have a problem with my server, its not responding anymore",
	)
	user2.MustReply(h, t, "Can you try to restart it?")
	// TODO: wipe channel before
	// TODO: multiple different problems with successful answers
	// TODO: import all messages into Chroma
}
