package base

import "github.com/funtimecoding/soil/pkg/chromium/protocol"

func (s *Stack) Body(identifier string) string {
	return protocol.NewIdentifier(s.Client, identifier).Body()
}
