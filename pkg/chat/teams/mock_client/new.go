package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/channel"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		channels:     map[string][]*channel.Channel{},
		messages:     map[string][]*message.Message{},
		replies:      map[string][]*message.Message{},
		chatMessages: map[string][]*message.Message{},
		chatMembers:  map[string][]*member.Member{},
	}
	t.Cleanup(result.Close)

	return result
}
