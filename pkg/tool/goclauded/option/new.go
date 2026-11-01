package option

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/transcript_cache"
	memoryClient "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client"
	memoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	queryClient "github.com/funtimecoding/soil/pkg/tool/goqueryd/client"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"time"
)

func New() *Option {
	return &Option{
		Clock: func() time.Time { return time.Now().UTC() },
		Claude: func() face.ClaudeSource {
			return transcript_cache.New(claude.New())
		},
		Memory:   func() memoryd.Client { return memoryClient.NewEnvironment() },
		Query:    func() queryd.Client { return queryClient.NewEnvironment() },
		Notifier: func() face.Notifier { return notifier.New() },
	}
}
