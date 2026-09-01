package scan

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system"

func (s *Service) checkVersionedPaths(
	v *virtual_file_system.System,
	path string,
) {
	s.Concerns = append(s.Concerns, versionedPathConcerns(v, path)...)
}
