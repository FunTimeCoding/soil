package page_put

import "github.com/funtimecoding/soil/pkg/notation"

func (p *Put) Encode() string {
	return notation.Encode(p, false)
}
