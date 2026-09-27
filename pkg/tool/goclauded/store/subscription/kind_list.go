package subscription

import (
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/split"
)

func (s *Subscription) KindList() []string {
	if s.Kinds == "" {
		return nil
	}

	return strings.DeleteEmpty(split.Comma(s.Kinds))
}
