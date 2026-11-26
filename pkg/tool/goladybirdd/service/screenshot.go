package service

import (
	"fmt"
	"os"
	"path/filepath"
)

func (s *Service) Screenshot(tab string) (string, error) {
	t, e := s.resolveTab(tab)

	if e != nil {
		return "", e
	}

	s.navigation.Settle(t)
	image, f := guarded(
		s,
		func() ([]byte, error) {
			return s.client.Screenshot(t)
		},
	)

	if f != nil {
		return "", f
	}

	path := filepath.Join(
		s.downloadDirectory,
		fmt.Sprintf("ladybird_%s.png", t),
	)

	if g := os.WriteFile(path, image, 0644); g != nil {
		return "", fmt.Errorf("write file: %w", g)
	}

	return path, nil
}
