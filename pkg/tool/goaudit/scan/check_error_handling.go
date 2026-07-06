package scan

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system"

func (s *Service) checkErrorHandling(
	v *virtual_file_system.System,
	path string,
	configuration *Configuration,
) {
	if !s.Generated || !s.Server {
		return
	}

	var suppress map[string][]string

	if configuration.Suppress != nil {
		suppress = configuration.Suppress[s.Name]
	}

	s.Concerns = append(
		s.Concerns,
		errorHandlingConcerns(v, path, s.Name, suppress)...,
	)
}
