package service

import (
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/raid/elite"
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
)

func (s *Service) EnrichFile(
	base string,
	name string,
) {
	if !strings.HasSuffix(name, constant.DetailedWvWKillSuffix) {
		return
	}

	var f elite.Fight
	notation.MustDecodeBytes(system.ReadBytes(base, name), &f, false)
	enriched, players := s.store.EnrichFight(name, &f)

	if !enriched {
		return
	}

	s.logger.Structured("enriched_fight", "file", name, "players", players)
}
