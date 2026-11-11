package bravo

import "example/pkg/alfa"

type Holder struct {
	Server *alfa.Server
}

func Build() *Holder {
	return &Holder{Server: alfa.New("bravo")}
}

func Pass() string {
	return describe(alfa.New("charlie"))
}

func describe(s *alfa.Server) string {
	return s.Name
}

func Plain() {
	alfa.New("delta")
}

func Make() *alfa.Server {
	return alfa.New("echo")
}

func Once(s *alfa.Server) *alfa.Server {
	return alfa.Wrap(s)
}

func Twice(s *alfa.Server) *alfa.Server {
	return alfa.Wrap(alfa.Wrap(s))
}
