package double

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func (s *Set) WriteSearchTranscript(identifier string) {
	system.WriteFile(
		filepath.Join(
			s.Harbor,
			join.Empty(identifier, soil.NotationLogExtension),
		),
		[]byte(
			`{"uuid":"u1","type":"user","timestamp":"2026-10-08T01:00:00Z","message":{"role":"user","content":"where is the legacy protocol pin"}}
{"uuid":"a1","type":"assistant","timestamp":"2026-10-08T01:01:00Z","message":{"role":"assistant","content":[{"type":"text","text":"the pin lives in setup.go"}]}}
{"uuid":"a2","type":"assistant","timestamp":"2026-10-08T01:02:00Z","message":{"role":"assistant","content":[{"type":"tool_use","id":"t1","name":"Bash","input":{"command":"go test ./..."}}]}}
`,
		),
		0644,
	)
}
