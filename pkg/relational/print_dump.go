package relational

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (d *Database) PrintDump() {
	fmt.Println(notation.MarshalIndent(d.Dump()))
}
