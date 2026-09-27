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
			reset := resetText(span.Parent().Next().Text())

			switch label {
			case constant.UsageMeterSession:
				result.SessionPercent = percent
				result.SessionReset = reset
			case constant.UsageMeterWeek:
				result.WeeklyAllPercent = percent
				result.WeeklyAllReset = reset
			case constant.UsageMeterFableWeek:
				result.FablePercent = percent
				result.FableReset = reset
			default:
				return
			}

			found[label] = true
		},
	)

	if !found[constant.UsageMeterSession] ||
		!found[constant.UsageMeterWeek] ||
		!found[constant.UsageMeterFableWeek] {
		return nil
	}

	return result
}
