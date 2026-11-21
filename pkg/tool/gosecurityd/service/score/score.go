package score

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/score"

func (s *Service) Score(identifier string) (*score.Score, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	result, okay := s.scores[identifier]

	return result, okay
}
