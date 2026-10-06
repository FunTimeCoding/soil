package stamp

import "fmt"

func (s *Stamp) Text() string {
	return fmt.Sprintf(
		"Version: %s\nGitHash: %s\nCommitDate: %s\nModule: %s\nDirty: %t\n",
		s.DisplayVersion(),
		s.GitHash,
		s.CommitDate,
		s.Module,
		s.Dirty,
	)
}
