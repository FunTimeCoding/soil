package layout

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board/layout/entry"

type Section struct {
	Name    string         `yaml:"name"`
	Entries []*entry.Entry `yaml:"entries"`
}
