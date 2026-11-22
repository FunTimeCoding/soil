package feed

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/probe"
)

func (s *Service) store(
	host string,
	output string,
) error {
	probed, e := probe.Parse(output)

	if e != nil {
		return e
	}

	snapshot, f := convert.ProbeSnapshot(probed)

	if f != nil {
		return f
	}

	return s.system.Put(convert.Snapshot(host, snapshot))
}
