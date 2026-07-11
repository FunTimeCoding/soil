package service

import "go/token"

type qualifiedPosition struct {
	position token.Position
	oldName  string
	newName  string
}
