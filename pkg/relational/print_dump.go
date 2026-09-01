package relational

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (d *Database) PrintDump() {
	console.Line(notation.MarshalIndent(d.Dump()))
}
