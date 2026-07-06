package sentence

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Sentence) Join() string {
	count := s.Count()

	if count == 0 {
		return fmt.Sprintf("%s.", s.action)
	}

	result := []string{s.action}
	result = append(result, s.affect...)
	last := len(result) - 1
	result[last] = fmt.Sprintf("and %s", result[last])

	return fmt.Sprintf("%s.", join.CommaSpace(result))
}
