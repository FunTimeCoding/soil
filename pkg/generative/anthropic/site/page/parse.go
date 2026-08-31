package page

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/hypertext"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strconv"
	"strings"
)

func Parse(markup string) *Usage {
	document := hypertext.Document(strings.NewReader(markup))
	result := New(0, "", 0, "", 0, "")
	found := map[string]bool{}
	document.Find("div[role='meter']").Each(
		func(
			_ int,
			meter *goquery.Selection,
		) {
			span := document.Find(
				join.Empty("#", meter.AttrOr("aria-labelledby", "")),
			)
			percent, e := strconv.Atoi(meter.AttrOr("aria-valuenow", ""))

			if e != nil {
				return
			}

			label := strings.TrimSpace(span.Text())
			reset := strings.TrimPrefix(
				strings.TrimSpace(span.Parent().Next().Text()),
				constant.UsageResetPrefix,
			)

			switch label {
			case constant.UsageMeterSession:
				result.SessionPercent = percent
				result.SessionReset = reset
			case constant.UsageMeterAllModels:
				result.WeeklyAllPercent = percent
				result.WeeklyAllReset = reset
			case constant.UsageMeterFable:
				result.FablePercent = percent
				result.FableReset = reset
			default:
				return
			}

			found[label] = true
		},
	)

	if !found[constant.UsageMeterSession] ||
		!found[constant.UsageMeterAllModels] ||
		!found[constant.UsageMeterFable] {
		return nil
	}

	return result
}
