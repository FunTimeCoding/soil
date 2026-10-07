package runbook

import "github.com/funtimecoding/soil/pkg/text/markdown/runbook/section"

type Runbook struct {
	source   *[]byte
	Filename string
	Title    string
	Sections []*section.Section
}
