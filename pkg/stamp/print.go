package stamp

import "github.com/funtimecoding/soil/pkg/console"

func (s *Stamp) Print() {
	console.Format("%s", s.Text())
}
