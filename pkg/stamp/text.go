package stamp

import "fmt"

func (s *Stamp) Text() string {
	return fmt.Sprintf(
		"Version: %s\nGitHash: %s\nBuildDate: %s\nModule: %s\nDirty: %t\n",
		s.Version,
		s.GitHash,
		s.BuildDate,
		s.Module,
		s.Dirty,
	)
}
