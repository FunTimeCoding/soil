package service

import (
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/raid/elite"
	"os"
	"path/filepath"
	"strings"
)

func (s *Service) EnrichFile(
	base string,
	name string,
) error {
	if !strings.HasSuffix(name, constant.DetailedWvWKillSuffix) {
		return nil
	}

	b, e := os.ReadFile(filepath.Join(base, name))

	if e != nil {
		return e
	}

	var f elite.Fight

	if g := notation.DecodeBytes(b, &f); g != nil {
		return g
	}

	enriched, players := s.store.EnrichFight(name, &f)

	if !enriched {
		return nil
	}

	s.logger.Structured("enriched_fight", "file", name, "players", players)

	return nil
}
