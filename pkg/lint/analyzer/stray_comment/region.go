package stray_comment

import "go/token"

type Region struct {
	From token.Pos
	To   token.Pos
}
