package service

import "go/token"

type spliceEdit struct {
	start       token.Pos
	end         token.Pos
	replacement string
}
