package label_holder

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
)

func Apply(
	l face.Labeller,
	p *Plan,
	r *Rule,
	from string,
) (string, error) {
	var failures []error

	for _, u := range p.Unset {
		if _, e := l.ReleaseLabel(u.Identifier, r.Key, u.Value, from); e != nil {
			failures = append(failures, e)
		}
	}

	if p.Adopt == "" {
		return p.Holder, errors.Join(failures...)
	}

	claimed, e := l.ClaimLabel(p.Adopt, r.Key, r.Value, from)

	if e != nil {
		failures = append(failures, e)
	}

	if !claimed {
		return "", errors.Join(failures...)
	}

	return p.Holder, errors.Join(failures...)
}
