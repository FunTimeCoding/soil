package base

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/tab"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (s *Stack) tabs() []*tab.Tab {
	var result []*tab.Tab
	notation.MustDecode(
		web.GetString(
			s.pool.Client(),
			locator.New("localhost").
				Port(s.Port).
				Path(constant.ChromiumNotationPath).
				Insecure().
				String(),
		),
		&result,
		true,
	)

	return result
}
