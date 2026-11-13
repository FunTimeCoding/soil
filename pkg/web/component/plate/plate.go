package plate

import "github.com/funtimecoding/soil/pkg/web/component/fact/row"

type Plate struct {
	name     string
	headline string
	caption  string
	rows     []*row.Row
}
