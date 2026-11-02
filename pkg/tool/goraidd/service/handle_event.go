package service

import (
	"github.com/fsnotify/fsnotify"
	gw2Constant "github.com/funtimecoding/soil/pkg/gw2/constant"
	"path/filepath"
	"strings"
)

func (s *Service) HandleEvent(v fsnotify.Event) []error {
	if !v.Has(fsnotify.Create) && !v.Has(fsnotify.Write) {
		return nil
	}

	var result []error

	if filepath.Base(v.Name) == gw2Constant.LogFile {
		s.logger.Structured("log_cache_sync")
		result = append(result, s.SyncLogCache()...)
	}

	if strings.HasSuffix(v.Name, gw2Constant.DetailedWvWKillSuffix) {
		if e := s.EnrichFile(filepath.Dir(v.Name), filepath.Base(v.Name)); e != nil {
			result = append(result, e)
		}
	}

	return result
}
