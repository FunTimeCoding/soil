package score

import "github.com/funtimecoding/soil/pkg/vulnerability/national_database/score"

func New(v *score.Score) *Score {
	return &Score{
		Identifier: v.Identifier,
		Score:      v.Score,
		Severity:   v.Severity,
		Version:    v.Version,
		Modified:   v.Modified,
	}
}
