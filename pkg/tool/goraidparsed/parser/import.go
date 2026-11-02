package parser

import "path/filepath"

func (p *Parser) Import(
	template string,
	workdir string,
) error {
	return p.python(
		filepath.Join(p.path, "tiddler_import.py"),
		"--html",
		template,
		"--logfolder",
		workdir,
		"--download",
		workdir,
	)
}
