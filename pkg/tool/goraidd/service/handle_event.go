package service

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"path/filepath"
	"strings"
)

func (s *Service) HandleEvent(v fsnotify.Event) {
	if !v.Has(fsnotify.Create) && !v.Has(fsnotify.Write) {
		return
	}

	if filepath.Base(v.Name) == constant.LogFile {
		s.logger.Structured("log_cache_sync")
		s.SyncLogCache()
	}

	if strings.HasSuffix(v.Name, constant.DetailedWvWKillSuffix) {
		s.EnrichFile(filepath.Dir(v.Name), filepath.Base(v.Name))
	}
}
