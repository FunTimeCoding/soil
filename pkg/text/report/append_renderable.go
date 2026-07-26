package report

import "github.com/funtimecoding/soil/pkg/text/constant"

func (s *Section) appendRenderable(other renderable) {
	if s.maximumLength == constant.NoLimit ||
		s.Length()+other.Length() <= s.maximumLength {
		s.renderables = append(s.renderables, other)
	}
}
