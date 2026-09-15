package process

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/shirou/gopsutil/v4/process"
)

func (c *Client) Processes() []*Entry {
	list, e := process.Processes()
	errors.PanicOnError(e)
	var result []*Entry

	for _, p := range list {
		parent, f := p.Ppid()
		name, g := p.Name()

		if f != nil || g != nil {
			continue
		}

		result = append(result, NewEntry(p.Pid, parent, name))
	}

	return result
}
