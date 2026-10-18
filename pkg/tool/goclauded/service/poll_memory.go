package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/version_entry"
)

func (s *Service) PollMemory() {
	var versions []version_entry.Entry

	if s.memoryGuard.Do(
		context.Background(),
		func(x context.Context) error {
			var e error
			versions, e = s.memory.VersionsSince(x, s.lastMemoryPoll, 50)

			return e
		},
	) != nil {
		return
	}

	for _, v := range versions {
		kind := constant.QueueMemoryUpdate

		if v.ChangeType == "created" {
			kind = constant.QueueMemoryCreate
		}

		var body string

		if v.Source != "" {
			body = fmt.Sprintf("%s %s by %s", v.Name, v.ChangeType, v.Source)
		} else {
			body = fmt.Sprintf("%s %s", v.Name, v.ChangeType)
		}

		if e := s.PushQueueBroadcast(kind, body); e != nil {
			s.reporter.CaptureException(e)
		}
	}

	if len(versions) > 0 {
		s.lastMemoryPoll = s.clock().UTC().Format("2006-01-02T15:04:05Z")
	}
}
