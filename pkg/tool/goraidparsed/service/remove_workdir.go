package service

import "os"

func (s *Service) removeWorkdir(path string) {
	if e := os.RemoveAll(path); e != nil {
		s.logger.Structured("remove_workdir", "path", path, "error", e)
	}
}
