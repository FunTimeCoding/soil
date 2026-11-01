package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	memoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"time"
)

type Option struct {
	Address           string
	MetricAddress     string
	ServiceTokens     []string
	LitePath          string
	SessionExportPath string
	Harbor            string
	Clock             func() time.Time
	Claude            func() face.ClaudeSource
	Memory            func() memoryd.Client
	Query             func() queryd.Client
	Notifier          func() face.Notifier
}
