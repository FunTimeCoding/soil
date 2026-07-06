package reader

import "github.com/funtimecoding/soil/pkg/system"

func (r *Reader) Line() string {
	return system.ReadLine(r.reader)
}
