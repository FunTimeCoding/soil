package site

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (s *Site) Dump() {
	fmt.Println(s.protocol.Outer(constant.AnthropicBodyElement))
}
