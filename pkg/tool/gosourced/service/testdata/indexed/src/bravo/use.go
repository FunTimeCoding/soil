package bravo

import "example/alfa"

func Use() int {
	s := alfa.NewServer()
	s.Start()

	return s.Port
}
