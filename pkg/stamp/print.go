package stamp

import "github.com/funtimecoding/soil/pkg/console"

func (s *Stamp) Print() {
	console.Format(
		"Version: %s\nGitHash: %s\nBuildDate: %s\n",
		s.Version,
		s.GitHash,
		s.BuildDate,
	)
}
