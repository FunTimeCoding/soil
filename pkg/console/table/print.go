package table

import "fmt"

func (t *Table) Print() {
	fmt.Print(t.Render())
}
