package subscription

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"

func (s Subscription) Label() string {
	if s.Alias != "" {
		return s.Alias
	}

	r := []rune(s.RootIdentifier)

	if len(r) <= constant.LabelPrefix {
		return s.RootIdentifier
	}

	return string(r[:constant.LabelPrefix])
}
