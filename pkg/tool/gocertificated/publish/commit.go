package publish

import "gitlab.com/gitlab-org/api/client-go/v2"

func (p *Publisher) Commit(
	v []*Change,
	message string,
) (string, error) {
	var action []*gitlab.CommitActionOptions

	for _, c := range v {
		present, e := p.exists(c.Path)

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

	result, e := p.forge.CommitActions(p.project, p.branch, message, action)

	if e != nil {
		return "", e
	}

	return result.ID, nil
}
