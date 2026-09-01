package example

import (
	"github.com/funtimecoding/soil/pkg/assistant"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system"
)

func Listen() {
	a := assistant.NewEnvironment(
		assistant.WithSubscriber(
			func(m *message.Message) {
				if m.Event == nil {
					console.Format("non-event message: %s\n", m.Type)

					return
				}

				console.Format("%s %s\n", m.Type, m.Event.Origin)
				console.Format("  Time: %s\n", m.Event.Time)
				console.Format("  Raw: %s\n", string(m.Event.Raw))
				console.Line()
			},
		),
	)
	defer a.Stop()
	a.Start()
	system.KillSignalBlock()
}
