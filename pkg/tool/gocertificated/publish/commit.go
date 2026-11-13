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

	project, e := p.projectIdentifier(x)

	if e != nil {
		return "", e
	}

	result, f := p.forge.CommitActions(x, project, p.branch, message, action)

	if f != nil {
		return "", f
	}

	return result.Identifier, nil
}
