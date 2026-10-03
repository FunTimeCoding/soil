package report

import "github.com/funtimecoding/soil/pkg/stamp"

func New(
	name string,
	s *stamp.Stamp,
) *Report {
	return &Report{
		Name:      name,
		Version:   s.Version,
		GitHash:   s.GitHash,
		BuildDate: s.BuildDate,
		Module:    s.Module,
		Dirty:     s.Dirty,
	}
}
