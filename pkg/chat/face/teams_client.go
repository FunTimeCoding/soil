package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/channel"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/chat/teams/team"
)

type TeamsClient interface {
	ChannelMessages(
		team string,
		channel string,
		limit int,
	) ([]*message.Message, error)
	Channels(team string) ([]*channel.Channel, error)
	ChatMembers(chat string) ([]*member.Member, error)
	ChatMessages(
		chat string,
		limit int,
	) ([]*message.Message, error)
	Chats(limit int) ([]*chat.Chat, error)
	Close()
	Reachable(x context.Context) error
	Replies(
		team string,
		channel string,
		identifier string,
		limit int,
	) ([]*message.Message, error)
	SendChatMessage(
		chat string,
		text string,
		mentioned []*member.Member,
	) (*message.Message, error)
	Teams() ([]*team.Team, error)
}
