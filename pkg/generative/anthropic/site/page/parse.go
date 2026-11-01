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
	weekSeen := false
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

			reset := resetText(span.Parent().Next().Text())

			switch accessibleText(span) {
			case constant.UsageMeterSession:
				result.SessionPercent = percent
				result.SessionReset = reset
				result.SessionSeen = true
			case constant.UsageMeterWeek:
				result.WeeklyAllPercent = percent
				result.WeeklyAllReset = reset
				weekSeen = true
			case constant.UsageMeterFableWeek, constant.UsageMeterFable:
				result.FablePercent = percent
				result.FableReset = reset
				result.FableSeen = true
			}
		},
	)

	if !weekSeen {
		return nil
	}

	return result
}
