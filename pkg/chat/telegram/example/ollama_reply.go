package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/telegram"
	"github.com/funtimecoding/soil/pkg/chat/telegram/constant"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/text/multi_line"
)

func OllamaReply() {
	o := ollama.NewEnvironment()
	t := telegram.NewEnvironment()
	defer t.Close()
	f := constant.Format
	c := environment.Required(constant.ChannelEnvironment)
	l := multi_line.New()
	l.Format(
		"You are a helpful assistant called %s. Below is a set of lines from a chat channel. Respond with a message in a friendly and informative manner. I will worry about the formatting, time and your name. Try to blend in with the message length and writing style. You can also choose to say nothing if you feel it is not appropriate.",
		t.Self().UserName,
	)

	for _, m := range t.MessagesByChannel(c) {
		if false {
			fmt.Println(m.Format(f))
		}

		if m.Text == "" {
			continue
		}

		l.Add(m.FormatChat())
	}

	if l.Count() > 1 {
		g := o.GenerateFast(l.Render())

		if false {
			fmt.Println(g.Text)
			g.Print()
		}

		t.SendMessage(t.ChannelByName(c).Identifier, g.Text)
	} else {
		fmt.Println("Chat empty")
	}
}
