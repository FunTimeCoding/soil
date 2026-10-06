package board

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board/connection"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board/layout"
)

type Board struct {
	Connection connection.Connection `yaml:"connection"`
	Top        []*layout.Column      `yaml:"top"`
	Tail       layout.Tail           `yaml:"tail"`
}
