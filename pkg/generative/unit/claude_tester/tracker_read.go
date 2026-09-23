package claude_tester

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"

func TrackerRead(
	path string,
	s *tracker.State,
) error {
	_, e := tracker.Read(path, s, nil)

	return e
}
