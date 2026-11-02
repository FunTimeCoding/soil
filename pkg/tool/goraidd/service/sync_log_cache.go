package service

import (
	"github.com/funtimecoding/soil/pkg/gw2"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/log"
	"os"
	"path/filepath"
)

func (s *Service) SyncLogCache() []error {
	b, e := os.ReadFile(filepath.Join(s.logCachePath, constant.LogFile))

	if e != nil {
		return []error{e}
	}

	logs := log.NewSlice(gw2.ParseLogs(b, false))
	added := s.store.SaveLogEntries(logs)
	s.logger.Structured("log_cache_synced", "total", len(logs), "new", added)

	return nil
}
