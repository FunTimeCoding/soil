package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	githubConstant "github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/types/run_key"
)

func (s *Service) Latest() (map[run_key.Key]*run.Run, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (map[run_key.Key]*run.Run, error) {
			repositories, e := s.client.Repositories(s.owner)

			if not_found.Is(e) {
				return nil, not_configured.Format(
					constant.OwnerNotFound,
					s.owner,
				)
			}

			if e != nil {
				return nil, fmt.Errorf("list repositories: %w", e)
			}

			result := make(map[run_key.Key]*run.Run)

			for _, repository := range repositories {
				name := repository.GetName()
				runs, f := s.client.LatestRuns(s.owner, name)

				if f != nil {
					return nil, fmt.Errorf("list runs: %w", f)
				}

				for _, r := range runs {
					if r.Status != githubConstant.CompletedStatus {
						continue
					}

					k := run_key.Key{
						Repository: name,
						Workflow:   r.Name,
						Branch:     r.Branch,
					}

					if existing, okay := result[k]; !okay ||
						r.Identifier > existing.Identifier {
						result[k] = r
					}
				}
			}

			return result, nil
		},
	)
}
