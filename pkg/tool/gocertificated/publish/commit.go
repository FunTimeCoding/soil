package publish

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/change"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (p *Publisher) Commit(
	x context.Context,
	v []*change.Change,
	message string,
) (string, error) {
	var action []*gitlab.CommitActionOptions

	for _, c := range v {
		present, e := p.exists(x, c.Path)

		if e != nil {
			return "", e
		}

		kind := gitlab.FileCreate

		if present {
			kind = gitlab.FileUpdate
		}

		action = append(
			action,
			&gitlab.CommitActionOptions{
				Action:   &kind,
				FilePath: new(c.Path),
				Content:  new(c.Content),
			},
		)
	}

	result, e := p.forge.CommitActions(x, p.project, p.branch, message, action)

	if e != nil {
		return "", e
	}

	return result.Identifier, nil
}
