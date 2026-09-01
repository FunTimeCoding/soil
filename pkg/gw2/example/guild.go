package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gw2"
)

func Guild() {
	a := argument.NewSimple("gw2-guild")
	a.String(constant.Tag, "", "Guild tag")
	a.ParseSimple()
	c := gw2.NewEnvironment()
	tag := a.GetString(constant.Tag)
	account := c.Account()

	if len(account.GuildLeader) == 0 {
		console.Line("No guilds with leader permissions")

		return
	}

	if tag == "" {
		console.Format("No guild tag provided, not printing members\n")
	}

	var tagFound bool

	for _, l := range account.GuildLeader {
		g := c.Guild(l)
		console.Format("Guild: %s\n", g.Name)
		console.Format("  Tag: %s\n", g.Tag)

		if tag != "" && tag == g.Tag {
			tagFound = true
			members := c.Members(l)
			console.Format("  Members: %d\n", len(members))

			for _, member := range members {
				console.Format("    Member: %+v\n", member.Name)
			}
		}
	}

	if tag != "" && !tagFound {
		console.Format("Guild tag %s not found\n", tag)
	}
}
