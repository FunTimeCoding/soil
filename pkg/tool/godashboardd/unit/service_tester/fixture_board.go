package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
)

func FixtureBoard() *board.Board {
	return board.Load(fixture.Path(systemConstant.BoardPath, "board.yaml"))
}
