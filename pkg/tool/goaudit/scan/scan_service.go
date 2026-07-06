package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
)

func scanService(
	v *virtual_file_system.System,
	path string,
	name string,
	repo string,
	c *Configuration,
) *Service {
	s := &Service{
		Name: name,
		Repo: repo,
	}
	s.ModelContext = v.DirectoryExists(filepath.Join(path, "model_context"))
	s.Server = v.DirectoryExists(filepath.Join(path, "server"))
	s.Web = v.DirectoryExists(filepath.Join(path, "web"))
	s.Store = v.DirectoryExists(filepath.Join(path, "store"))
	s.Generated = v.DirectoryExists(filepath.Join(path, "generated"))
	s.Convert = v.DirectoryExists(filepath.Join(path, "convert"))
	s.Client = v.DirectoryExists(filepath.Join(path, "client"))
	s.Types = v.DirectoryExists(filepath.Join(path, "types"))
	s.Model = v.DirectoryExists(filepath.Join(path, "model"))
	s.ConstantDirectory = v.DirectoryExists(filepath.Join(path, "constant"))
	s.ConstantFile = !s.ConstantDirectory && v.Has(
		filepath.Join(
			path,
			"constant.go",
		),
	)
	s.Worker = v.DirectoryExists(filepath.Join(path, "worker"))
	s.IntegrationTests = v.DirectoryExists(
		filepath.Join(
			path,
			"integration_test",
		),
	)
	s.Option = v.DirectoryExists(filepath.Join(path, "option"))
	s.Run = v.Has(filepath.Join(path, "run.go"))

	if !s.hasCapability() {
		return nil
	}

	s.collectWarnings(v, path, c)

	return s
}
