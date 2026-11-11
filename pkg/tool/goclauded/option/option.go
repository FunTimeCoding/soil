package option

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/credential"
	generative "github.com/funtimecoding/soil/pkg/generative/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	memoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Option struct {
	option.Web
	Metric                 option.Web
	LitePath               string
	SessionExportDirectory string
	Harbor                 string
	MonitorUsage           bool
	UtilizationFallback    bool
	ChannelAlways          bool
	Clock                  func() time.Time
	Claude                 func() face.ClaudeSource
	Memory                 func() memoryd.Client
	Query                  func() queryd.Client
	Notifier               func() face.Notifier
	Browser                func() *chromium.Client
	Sources                func() []string
	Utilization            func() generative.AnthropicClient
	Credential             func() *credential.Credential
}
