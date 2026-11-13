package transcript

import (
	"encoding/json"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/memory_payload"
	gomemorydConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"strings"
)

func profileEntries(body string) (map[string][]memory_payload.Entry, bool) {
	var payload memory_payload.Profile

	if json.Unmarshal([]byte(body), &payload) == nil {
		return map[string][]memory_payload.Entry{
			constant.TierAlways:   payload.Always,
			constant.TierRelevant: payload.Relevant,
		}, true
	}

	lines := strings.Split(body, stringsConstant.Unix)
	always, alwaysReadable := textEntries(
		body,
		lines,
		gomemorydConstant.AlwaysSectionHeading,
	)
	relevant, relevantReadable := textEntries(
		body,
		lines,
		gomemorydConstant.RelevantSectionHeading,
	)

	return map[string][]memory_payload.Entry{
		constant.TierAlways:   always,
		constant.TierRelevant: relevant,
	}, alwaysReadable && relevantReadable
}
