package parser

import (
	"path/filepath"
	"time"
)

func (p *Parser) Parse(
	workdir string,
	date time.Time,
) error {
	return p.python(
		filepath.Join(p.path, "TW5_parse_top_stats_detailed.py"),
		workdir,
		"-d",
		date.Format("2006-01-02T15:04:05"),
	)
}
