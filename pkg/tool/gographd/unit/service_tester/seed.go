package service_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/channel"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/chat/teams/team"
)

func (o *Tester) Seed() {
	o.Source.AddTeam(team.New("team-1", "Security"))
	o.Source.AddTeam(team.New("team-2", "Operations"))
	o.Source.AddChannel("team-1", channel.New("channel-1", "General"))
	o.Source.AddChannel("team-1", channel.New("channel-2", "CVE-2026-1"))
	root := message.New("100", "Norbert", "patched?")
	root.Replies = 1
	o.Source.AddMessage("team-1", "channel-2", root)
	o.Source.AddReply(
		"team-1",
		"channel-2",
		"100",
		message.New("101", "Alexander", "yes"),
	)
	squad := chat.New("19:alpha@thread.v2", "group")
	squad.Topic = "Squad"
	o.Source.AddChat(squad)
	o.Source.AddChatMessage(
		"19:alpha@thread.v2",
		message.New("200", "Knut", "status?"),
	)
}
