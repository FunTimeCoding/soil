package site

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
)

func (s *Site) printCloseMemories() {
	s.protocol.PrintNode(CloseMemoriesSelector, usefulAttributes)
	n := s.protocol.Select(CloseMemoriesSelector, 2)
	fmt.Println("Close dialog index 2")
	protocol.Print(n, usefulAttributes)
}
