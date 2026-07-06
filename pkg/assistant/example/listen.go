package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assistant"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/system"
)

func Listen() {
	a := assistant.NewEnvironment(
		assistant.WithSubscriber(
			func(m *message.Message) {
				if m.Event == nil {
					fmt.Printf("non-event message: %s\n", m.Type)

					return
				}

				fmt.Printf("%s %s\n", m.Type, m.Event.Origin)
				fmt.Printf("  Time: %s\n", m.Event.Time)
				fmt.Printf("  Raw: %s\n", string(m.Event.Raw))
				fmt.Println()
			},
		),
	)
	defer a.Stop()
	a.Start()
	system.KillSignalBlock()
}
