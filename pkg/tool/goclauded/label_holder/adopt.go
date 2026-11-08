package label_holder

import "github.com/funtimecoding/soil/pkg/tool/goclauded/face"

func adopt(
	l face.Labeller,
	sessionIdentifier string,
	r *Rule,
	from string,
) (bool, error) {
	if len(r.Compatible) == 0 {
		return l.ClaimLabel(sessionIdentifier, r.Key, r.Value, from)
	}

	if e := l.AddLabel(sessionIdentifier, r.Key, r.Value, from); e != nil {
		return false, e
	}

	return true, nil
}
