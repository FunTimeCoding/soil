package channel

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel/model_context_sink"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/google/uuid"
	"time"
)

func New(
	i *identity.Tool,
	version string,
	instructions string,
) *Server {
	result := &Server{
		nonce: uuid.New().String(),
		sleep: time.Sleep,
		ready: make(chan struct{}),
		open:  make(chan struct{}),
	}
	result.server = server.New(i, version).
		WithInstructions(instructions).
		WithExperimental(
			map[string]any{constant.ChannelCapability: map[string]any{}},
		).
		WithConnected(result.Connected).
		Server()
	result.sink = model_context_sink.New(result.server)
	result.register()

	return result
}
