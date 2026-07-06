package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/telegram"
	"github.com/funtimecoding/soil/pkg/chat/telegram/constant"
	"github.com/funtimecoding/soil/pkg/chat/telegram/message"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant/prompts"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/text/multi_line"
	timeLibrary "github.com/funtimecoding/soil/pkg/time"
	"strings"
	"time"
)

func OllamaSession() {
	o := ollama.NewEnvironment()
	t := telegram.NewEnvironment()
	defer t.Close()
	t.PrintDatabase()
	f := constant.Format
	c := environment.Required(constant.ChannelEnvironment)
	verbose := true
	statistics := false
	ownName := t.Self().UserName
	var messages []*message.Message

	if false {
		messages = t.MessagesByChannel(c)

		if len(messages) > 0 {
			fmt.Println("Messages already in channel")

			for _, m := range messages {
				fmt.Println(m.Format(f))
			}
		}
	}

	go func() {
		fmt.Printf("Listening to channel %s\n", c)

		for u := range t.Listen() {
			if u.FromChat().Title != c {
				continue
			}

			m := message.New(u.Message)
			messages = append(messages, m)

			if u.SentFrom().UserName == ownName {
				continue
			}

			fmt.Println(m.Format(f))
			history := multi_line.New()

			for _, e := range messages {
				if e.Text == "" {
					continue
				}

				history.Add(e.FormatChat())
			}

			if history.Count() == 0 {
				fmt.Println("Chat empty")

				continue
			}

			p := prompts.DecideAction()
			p.History(history.Render())
			decision := o.GenerateSimple(p.Render())

			if verbose {
				fmt.Printf("Decision: %s\n", decision.Text)

				if statistics {
					decision.Print()
				}
			}

			if strings.Contains(
				strings.ToLower(decision.Text),
				"no action needed",
			) {
				continue
			}

			l := multi_line.New()
			l.Format(
				"You are a LLM-driven ChatOps bot called %s. Below is the current snapshot of messages from your Telegram chat channel %s. Do not re-digest the whole log every time by replying to earlier messages. Blend in and respond with a message in a informative manner. Do not prefix your response with a timestamp and your name. Try to match the other users message length and writing style. The time is %s, so you know how old previous messages are.",
				ownName,
				c,
				time.Now().Format(timeLibrary.DateMinute),
			)
			l.Add(history.Render())
			response := o.GenerateSimple(l.Render())

			if verbose {
				fmt.Printf("Response: %s\n", response.Text)

				if statistics {
					response.Print()
				}
			}

			t.SendMessage(t.ChannelByName(c).Identifier, response.Text)
		}
	}()
	system.KillSignalBlock()
	t.Close()
}
