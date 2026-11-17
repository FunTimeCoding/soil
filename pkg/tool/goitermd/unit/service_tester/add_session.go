package service_tester

import "github.com/funtimecoding/soil/pkg/iterm/session"

func (o *Tester) AddSession(identifier string) {
	s := session.Stub()
	s.Identifier = identifier
	o.Source.AddSession(s)
}
