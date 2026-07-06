package main

import (
	"github.com/funtimecoding/soil/pkg/bubbletea"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/example_list"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/example_table"
	"github.com/funtimecoding/soil/pkg/bubbletea/table/example_country"
)

func main() {
	if false {
		bubbletea.Run(example_list.New(), false)
	}

	if true {
		bubbletea.Run(example_table.New(example_country.New()), false)
	}
}
