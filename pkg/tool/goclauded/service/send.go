package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Service) Send(
	name string,
	to string,
	body string,
	immediate bool,
) (bool, error) {
	var holder *session.Session

	if to != "" {
		found, e := s.store.SessionByCallsign(to)

		if e != nil {
			return false, e
		}

		if found == nil {
			return false, not_found.New(constant.Callsign, to)
		}

		holder = found
	}

	m, e := s.store.SendMessage(name, to, body)

	if e != nil {
		return false, e
	}

	formatted := fmt.Sprintf("%s: %s", name, body)

	if holder == nil {
		template := queue.NewBroadcast(constant.QueueMessage, formatted)
		template.MessageIdentifier = new(m.Identifier)

		return false, s.pushEntryBroadcast(template)
	}

	entry := queue.New(holder.Identifier, to, constant.QueueMessage, formatted)
	entry.MessageIdentifier = new(m.Identifier)

	if immediate {
		return s.pushImmediate(entry)
	}

	return false, s.pushEntry(entry)
}
