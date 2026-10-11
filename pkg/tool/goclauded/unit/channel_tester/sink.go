package channel_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel/mock_sink"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
)

func Sink() (*channel.Server, *mock_sink.Sink) {
	k := mock_sink.New()

	return channel.New(constant.Identity, "").WithSink(k), k
}
