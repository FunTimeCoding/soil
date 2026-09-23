package service

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	chromiumConstant "github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"strings"
)

func usageTabEvidence() (result []any) {
	defer func() {
		if recover() != nil {
			result = []any{"browser", "unreachable"}
		}
	}()
	c := chromium.NewEnvironment()
	defer c.Close()
	var count int
	var anthropic []string

	for _, t := range c.Tabs() {
		if t.Type != chromiumConstant.PageTabType {
			continue
		}

		count++

		if strings.Contains(t.Locator, constant.AnthropicSiteHost) {
			anthropic = append(anthropic, t.Locator)
		}
	}

	return []any{"page_count", count, "anthropic_locator", anthropic}
}
