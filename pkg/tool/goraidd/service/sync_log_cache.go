package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gw2"
	gw2Constant "github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/log"
	"os"
	"path/filepath"
)

func (s *Service) SyncLogCache(x context.Context) []error {
	b, e := os.ReadFile(filepath.Join(s.logCachePath, gw2Constant.LogFile))

	if e != nil {
		return []error{e}
	}

	if x.Err() != nil {
		return nil
	}

	logs := log.NewSlice(gw2.ParseLogs(b, false))
	added := s.store.SaveLogEntries(logs)
	s.logger.Structured("log_cache_synced", "total", len(logs), "new", added)

	return nil
}
