package runbook

import "github.com/yuin/goldmark/v2/parser"

func (r *Runbook) Parse(filename string) {
	o := parser.New().Parse(*r.source)
	r.Filename = filename
	r.Walk(o)
}
