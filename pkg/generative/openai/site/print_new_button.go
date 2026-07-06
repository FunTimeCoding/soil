package site

import "github.com/funtimecoding/soil/pkg/chromium/protocol"

func (s *Site) printNewButton() {
	s.protocol.PrintNode(NewSelector, usefulAttributes)
	protocol.Print(
		s.protocol.Select(NewSelector, 0),
		usefulAttributes,
	)
}
