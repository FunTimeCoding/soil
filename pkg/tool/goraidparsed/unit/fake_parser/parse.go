package fake_parser

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"time"
)

func (p *Parser) Parse(
	workdir string,
	_ time.Time,
) error {
	p.calls++
	p.workdir = workdir

	if p.failure != nil {
		return p.failure
	}

	for i := range p.tiddlers {
		system.SaveFile(
			filepath.Join(workdir, fmt.Sprintf("fight-%d.tid", i)),
			"tiddler",
		)
	}

	return nil
}
