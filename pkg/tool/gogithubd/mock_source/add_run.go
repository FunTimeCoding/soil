package mock_source

import "github.com/funtimecoding/soil/pkg/github/run"

func (s *Source) AddRun(
	repository string,
	r *run.Run,
) {
	s.runs[repository] = append(s.runs[repository], r)
}
