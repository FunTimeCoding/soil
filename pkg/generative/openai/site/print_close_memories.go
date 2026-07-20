package site

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/generative/openai/constant"
)

func (s *Site) printCloseMemories() {
	s.protocol.PrintNode(constant.CloseMemoriesSelector, constant.UsefulAttributes)
	n := s.protocol.Select(constant.CloseMemoriesSelector, 2)
	fmt.Println("Close dialog index 2")
	protocol.Print(n, constant.UsefulAttributes)
}
