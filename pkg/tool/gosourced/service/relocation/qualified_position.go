package relocation

import "go/token"

type QualifiedPosition struct {
	Position token.Position
	OldName  string
	NewName  string
}
