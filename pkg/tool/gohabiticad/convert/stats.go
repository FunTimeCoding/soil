package convert

import (
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func Stats(s *statistic.Statistic) *server.Stats {
	return &server.Stats{
		Hp:     float32(s.HP),
		Mp:     float32(s.MP),
		Xp:     float32(s.XP),
		Gp:     float32(s.GP),
		Level:  s.Level,
		Class:  s.Class,
		Points: s.Points,
		Str:    s.Str,
		Con:    s.Con,
		Int:    s.Int,
		Per:    s.Per,
	}
}
