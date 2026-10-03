package site

import (
	"github.com/funtimecoding/soil/pkg/console"
	"time"
)

func (s *Site) ExtractFlow(verbose bool) string {
	s.NewChat()
	s.clickProfile()
	s.clickSettings()
	s.clickPersonalize()
	s.clickMemories()
	time.Sleep(2 * time.Second)
	result := s.readMemories()

	if verbose {
		console.Format("Memories: %d\n", len(result))
	}

	s.clickCloseMemories()
	s.clickCloseSettings()

	return result
}
