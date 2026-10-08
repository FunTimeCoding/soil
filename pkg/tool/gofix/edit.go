package gofix

import "go/token"

type Edit struct {
	position token.Pos
	end      token.Pos
	newText  string
}
