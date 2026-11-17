package peer_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"path/filepath"
	"time"
)

func (s *Peers) DaemonOption() *option.Option {
	s.t.Helper()

	return s.Option(
		filepath.Join(s.t.TempDir(), "goclauded.sqlite"),
		func() time.Time { return time.Now().UTC() },
	)
}
