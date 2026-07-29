package silence

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (s *Silence) IsPermanent() bool {
	return s.CommentContains(constant.PermanentTag)
}
