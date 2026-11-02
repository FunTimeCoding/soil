package transcript

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/sweep"
)

func (s *Service) SyncChanged(
	identifier string,
	sourcePath string,
) {
	if sourcePath != "" && sweep.SyncFile(sourcePath, s.harbor) {
		s.logger.Structured(
			"source_file_shrunk",
			constant.Identifier,
			identifier,
		)
	}

	s.EnrichSession(identifier)
}
