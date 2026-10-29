package mock_source

import "github.com/funtimecoding/soil/pkg/github/run"

func (s *Source) LatestRuns(
	_ string,
	repository string,
) ([]*run.Run, error) {
	return s.runs[repository], nil
}
