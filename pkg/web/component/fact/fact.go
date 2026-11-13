package fact

import "github.com/funtimecoding/soil/pkg/web/component/fact/row"

type Fact struct {
	rows    []*row.Row
	columns int
	glance  bool
}
