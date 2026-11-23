package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
)

func (c *Client) chosen(
	name string,
	s *system_selection.Selection,
) bool {
	v := system_result.NewSummary(system.New(name, "", 0))

	if d := c.detail(name); d != nil {
		v.Coverage = d.Summary.Coverage
		v.TopScore = d.Summary.TopScore

		if d.Summary.Labels != nil {
			v.Labels = *d.Summary.Labels
		}
	}

	return s.Matches(v)
}
