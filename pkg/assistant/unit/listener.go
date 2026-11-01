package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assistant"
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/funtimecoding/soil/pkg/assistant/unit/fake_assistant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"time"
)

func listener(
	f *fake_assistant.Server,
	r *memory.Memory,
	received chan *message.Message,
) *assistant.Client {
	return assistant.New(
		f.Host(),
		"alfa",
		logger.New(context.Background()),
		r,
		assistant.WithSubscriber(
			func(m *message.Message) {
				if m.Event != nil {
					received <- m
				}
			},
		),
	).WithBackoff(time.Millisecond, 4*time.Millisecond)
}
