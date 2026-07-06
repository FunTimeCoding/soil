package scan

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system"

func (s *Service) collectWarnings(
	v *virtual_file_system.System,
	path string,
	configuration *Configuration,
) {
	s.checkStaleDirectories(v, path)
	s.checkModelContext(v, path)
	s.checkTopLevelArgument(v, path)
	s.checkTopLevelResponse(v, path)
	s.checkConstantFile(path)
	s.checkOption(path)
	s.checkRun(path)
	s.checkSuffix(path)
	s.checkIdentity(v, path)
	s.checkOpenAPI(v, path)
	s.checkStrictServer(v, path)
	s.checkErrorHandling(v, path, configuration)
	s.checkServerPatterns(v, path)
}
