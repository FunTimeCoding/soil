package service

import "github.com/funtimecoding/soil/pkg/system/run"

func (s *Service) python(arguments ...string) error {
	r := run.New().NoPanic()
	r.Directory = s.parserPath
	r.Start(append([]string{"python"}, arguments...)...)

	return r.Error
}
