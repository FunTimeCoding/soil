package service

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/receipt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/sweep"
	"os"
	"path/filepath"
)

func (s *Service) DeleteSession(
	identifier string,
	confirm string,
) (*receipt.Receipt, error) {
	r, e := s.deleteGate(identifier, confirm)

	if e != nil {
		return nil, e
	}

	result, f := s.deleteCounts(r)

	if f != nil {
		return nil, f
	}

	transcript := filepath.Join(
		s.harbor,
		join.Empty(identifier, constant.NotationLogExtension),
	)

	if _, g := os.Stat(transcript); g == nil {
		result.Transcript = transcript
	}

	if g := s.deleteLifetimeWindow(r, result); g != nil {
		return nil, g
	}

	s.client.Delete(identifier)
	result.Sources = sweep.DeleteSource(identifier)

	if g := s.store.DeleteSession(identifier); g != nil {
		return nil, g
	}

	if g := s.store.DeleteTrackerState(identifier); g != nil {
		return nil, g
	}

	s.cache.Delete(identifier)
	s.notify()

	return result, nil
}
