package channel_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel/mock_sink"
	goclaudeConstant "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
)

func Sink() (*channel.Server, *mock_sink.Sink) {
	k := mock_sink.New()

	return channel.New(goclaudeConstant.Identity, "", memory.New()).
		WithSink(k), k
}
