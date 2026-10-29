package service

import (
	"github.com/funtimecoding/soil/pkg/gw2"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/soil/pkg/system"
)

func (s *Service) SyncLogCache() {
	logs := log.NewSlice(
		gw2.ParseLogs(system.ReadBytes(s.logCachePath, constant.LogFile), false),
	)
	added := s.store.SaveLogEntries(logs)
	s.logger.Structured("log_cache_synced", "total", len(logs), "new", added)
}
