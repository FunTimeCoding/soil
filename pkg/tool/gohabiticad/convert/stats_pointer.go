package convert

import (
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func statsPointer(s *statistic.Statistic) *server.Stats {
	return Stats(s)
}
