package score

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/mirror"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/score"
	nvdScore "github.com/funtimecoding/soil/pkg/vulnerability/national_database/score"
)

func (s *Service) apply(
	page []*nvdScore.Score,
	m *mirror.Mirror,
) error {
	rows := make([]*score.Score, 0, len(page))

	for _, v := range page {
		rows = append(rows, score.New(v))
	}

	if e := s.store.PutScores(rows, m); e != nil {
		return e
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, r := range rows {
		s.scores[r.Identifier] = r
	}

	copied := *m
	s.state = &copied

	if len(rows) > 0 {
		s.generation++
	}

	return nil
}
