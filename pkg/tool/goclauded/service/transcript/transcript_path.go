package transcript

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
)

func (s *Service) transcriptPath(identifier string) string {
	return filepath.Join(
		s.harbor,
		join.Empty(identifier, soil.NotationLogExtension),
	)
}
