package service_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/channel"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/chat/teams/team"
)

func (o *Tester) Seed() {
	o.Source.AddTeam(team.New("team-1", "Kitchen"))
	o.Source.AddTeam(team.New("team-2", "Garden"))
	o.Source.AddChannel("team-1", channel.New("channel-1", "General"))
	o.Source.AddChannel("team-1", channel.New("channel-2", "Lunch"))
	root := message.New("100", "November", "lunch at noon?")
	root.Replies = 1
	o.Source.AddMessage("team-1", "channel-2", root)
	o.Source.AddReply(
		"team-1",
		"channel-2",
		"100",
		message.New("101", "Juliett", "yes"),
	)
	club := chat.New("19:alpha@thread.v2", "group")
	club.Topic = "Book club"
	o.Source.AddChat(club)
	o.Source.AddChatMessage(
		"19:alpha@thread.v2",
		message.New("200", "Kilo", "status?"),
	)
}
