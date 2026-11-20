package grant_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/grant"
)

func (r *Tester) Check(
	directory string,
	command string,
) (string, []string) {
	r.t.Helper()
	g := grant.New(r.Mock, r.Session, directory)

	return guard.GrantedBlockReason("darwin", command, g), g.Notes()
}
