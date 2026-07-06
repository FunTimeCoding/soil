package bubbletea

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/errors"
)

func RunProgram(p *tea.Program) {
	_, e := p.Run()
	errors.PanicOnError(e)
}
