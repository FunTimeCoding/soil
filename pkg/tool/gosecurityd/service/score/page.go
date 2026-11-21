package score

import (
	"context"
	"github.com/funtimecoding/soil/pkg/vulnerability/national_database/score_page"
	"time"
)

func (s *Service) page(
	x context.Context,
	start int,
	from time.Time,
	to time.Time,
) (*score_page.Page, error) {
	var result *score_page.Page
	e := s.guard.Attempt(
		x,
		func(y context.Context) error {
			var f error
			result, f = s.client.Scores(y, start, from, to)

			return f
		},
	)

	return result, e
}
