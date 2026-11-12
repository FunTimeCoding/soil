package entry

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board/layout/row"

type Entry struct {
	Label  string     `yaml:"label"`
	Link   string     `yaml:"link"`
	Icon   string     `yaml:"icon"`
	Widget string     `yaml:"widget"`
	Rows   []*row.Row `yaml:"rows"`
}
