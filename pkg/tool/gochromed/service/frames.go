package service

import (
	"context"
	chromiumConstant "github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) frames(identifier string) []*tab.Tab {
	tabs, e := peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*tab.Tab, error) { return s.source.Tabs() },
	)

	if e != nil {
		return nil
	}

	var result []*tab.Tab

	for _, t := range tabs {
		if t.Type == chromiumConstant.IframeTabType &&
			t.ParentIdentifier == identifier {
			result = append(result, t)
		}
	}

	return result
}
