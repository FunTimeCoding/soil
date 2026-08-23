package stamp

import "fmt"

func (s *Stamp) Print() {
	fmt.Printf(
		"Version: %s\nGitHash: %s\nBuildDate: %s\n",
		s.Version,
		s.GitHash,
		s.BuildDate,
	)
}
