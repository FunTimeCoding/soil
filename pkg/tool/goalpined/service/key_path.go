package service

import "path/filepath"

func (s *Service) keyPath() string {
	return filepath.Join(s.keyDirectory, s.signatureKey)
}
