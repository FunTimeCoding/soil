package score

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/score_result"

func (s *Service) Status() *score_result.Result {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return score_result.New(
		s.state.Filled,
		len(s.scores),
		s.state.Start,
		s.state.Total,
		s.state.Synced,
	)
}
