package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/dashboard_view"
)

func (s *Service) countMemories(v *dashboard_view.View) error {
	memories, e := s.store.ListMemories("", "", constant.AllScope, true)

	if e != nil {
		return e
	}

	tags := map[string]int{}

	for _, m := range memories {
		if s.isHidden(m.Tags) {
			continue
		}

		v.Active++
		v.Types[m.Type]++

		for _, t := range m.Tags {
			tags[t]++
		}
	}

	v.Tags = tagCounts(tags)

	return nil
}
