package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/channel"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/chat/teams/team"
)

type Client struct {
	Failure      error
	Closed       int
	teams        []*team.Team
	channels     map[string][]*channel.Channel
	messages     map[string][]*message.Message
	replies      map[string][]*message.Message
	chats        []*chat.Chat
	chatMessages map[string][]*message.Message
}
