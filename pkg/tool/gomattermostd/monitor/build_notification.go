package monitor

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"strings"
	"time"
)

func buildNotification(
	posts []*post.Post,
	channelName, username string,
) string {
	if len(posts) == 0 {
		return ""
	}

	var sb strings.Builder
	writer.Print(
		&sb,
		"@%s I found discussion about topics you're interested in!\n\n",
		username,
	)
	writer.Print(&sb, "**Channel:** %s\n\n", channelName)

	for i, p := range posts {
		writer.Print(
			&sb,
			"**Message %d** (%s):\n%s\n\n",
			i+1,
			p.Create.Format(time.RFC1123),
			p.Message,
		)
	}

	return sb.String()
}
