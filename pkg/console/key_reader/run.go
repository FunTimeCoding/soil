package key_reader

import (
	"github.com/funtimecoding/soil/pkg/bubbletea"
	"github.com/funtimecoding/soil/pkg/console/key_reader/model"
)

func (r *Reader) Run() {
	bubbletea.Run(model.New(r), false)
}
