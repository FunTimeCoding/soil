package report

import "github.com/funtimecoding/soil/pkg/console"

func (r *Report) Print() {
	console.Line(r.Encode())
}
