package call_format

import "go/token"

type Fix struct {
	Position token.Pos
	End      token.Pos
	NewText  string
}
