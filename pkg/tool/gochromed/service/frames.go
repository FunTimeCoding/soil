package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/browser/chromium/tab"
	browserConstant "github.com/funtimecoding/soil/pkg/browser/constant"
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
		if t.Type == browserConstant.ChromiumIframeTabType &&
			t.ParentIdentifier == identifier {
			result = append(result, t)
		}
	}

	return result
}
