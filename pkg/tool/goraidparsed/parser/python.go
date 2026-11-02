package parser

import "github.com/funtimecoding/soil/pkg/system/run"

func (p *Parser) python(arguments ...string) error {
	r := run.New().NoPanic()
	r.Directory = p.path
	r.Start(append([]string{"python"}, arguments...)...)

	return r.Error
}
