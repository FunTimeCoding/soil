package site

import (
	"github.com/funtimecoding/soil/pkg/console"
	"time"
)

func (s *Site) ExtractFlow(verbose bool) string {
	s.NewChat()

	if false {
		s.printProfile()
	}

	s.clickProfile()
	s.clickSettings()
	s.clickPersonalize()

	if false {
		s.printMemories()
	}

	s.clickMemories()
	time.Sleep(2 * time.Second)
	result := s.readMemories()

	if verbose {
		console.Format("Memories: %d\n", len(result))
	}

	if false {
		// Unstable selector
		s.printCloseMemories()
	}

	s.clickCloseMemories()

	if false {
		s.printCloseSettings()
	}

	s.clickCloseSettings()

	return result
}
