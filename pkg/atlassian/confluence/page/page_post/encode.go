package page_post

import "github.com/funtimecoding/soil/pkg/notation"

func (p *Post) Encode() string {
	return notation.Encode(p, false)
}
