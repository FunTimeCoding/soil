package change

import "github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"

func (c *Change) Apply(e *entry.Entry) {
	if c.Action != nil {
		e.Action = *c.Action
	}

	if c.User != nil {
		e.User = *c.User
	}

	if c.System != nil {
		e.System = *c.System
	}

	if c.Service != nil {
		e.Service = *c.Service
	}

	if c.Description != nil {
		e.Description = *c.Description
	}

	if c.Timestamp != nil {
		e.Timestamp = *c.Timestamp
	}
}
