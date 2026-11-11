package double

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	memoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"time"
)

func (s *Set) Option(
	litePath string,
	clock func() time.Time,
) *option.Option {
	result := option.New()
	result.LitePath = litePath
	result.SessionExportDirectory = s.t.TempDir()
	result.Harbor = s.Harbor
	result.Clock = clock
	result.Claude = func() face.ClaudeSource { return s.Client }
	result.Memory = func() memoryd.Client { return s.MemoryPeer }
	result.Query = func() queryd.Client { return s.Query }
	result.Notifier = func() soilFace.EventNotifier { return s.Notifier }
	result.Sources = func() []string { return []string{s.Source} }
	result.ChannelAlways = s.ChannelAlways

	return result
}
