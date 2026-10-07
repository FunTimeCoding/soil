package process

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/types/process_entry"
	"github.com/shirou/gopsutil/v4/process"
)

func (c *Client) Processes() []*process_entry.Entry {
	list, e := process.Processes()
	errors.PanicOnError(e)
	var result []*process_entry.Entry

	for _, p := range list {
		parent, f := p.Ppid()
		name, g := p.Name()

		if f != nil || g != nil {
			continue
		}

		result = append(result, process_entry.New(p.Pid, parent, name))
	}

	return result
}
