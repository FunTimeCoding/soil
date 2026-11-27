package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/channel"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/chat/teams/destination"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/chat/teams/team"
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
)

type TeamsClient interface {
	ChannelMessages(
		team string,
		channel string,
		limit int,
	) ([]*message.Message, error)
	Channels(team string) ([]*channel.Channel, error)
	ChatMessages(
		chat string,
		limit int,
	) ([]*message.Message, error)
	Chats(limit int) ([]*chat.Chat, error)
	Close()
	CreateChat(
		kind string,
		users []string,
		topic string,
	) (*chat.Chat, error)
	Me() (*user.User, error)
	Members(d *destination.Destination) ([]*member.Member, error)
	Post(
		d *destination.Destination,
		text string,
		mentioned []*member.Member,
	) (*message.Message, error)
	Reachable(x context.Context) error
	Replies(
		team string,
		channel string,
		identifier string,
		limit int,
	) ([]*message.Message, error)
	Teams() ([]*team.Team, error)
	User(identifier string) (*user.User, error)
	Users(
		query string,
		limit int,
	) ([]*user.User, error)
}
