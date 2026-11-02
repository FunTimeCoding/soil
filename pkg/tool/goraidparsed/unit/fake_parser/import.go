package fake_parser

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func (p *Parser) Import(
	template string,
	workdir string,
) error {
	p.calls++
	system.SaveFile(filepath.Join(workdir, filepath.Base(template)), "template")
	system.SaveFile(filepath.Join(workdir, "report.html"), "report")

	return nil
}
