package assert_call

import "go/token"

func (r Range) Contains(p token.Pos) bool {
	return p >= r.Start && p < r.End
}
