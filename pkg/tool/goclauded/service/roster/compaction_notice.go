package roster

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/shorten"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func compactionNotice(
	callsign string,
	phase constant.CompactionPhase,
	trigger string,
	topic string,
) string {
	result := fmt.Sprintf(constant.CompactionNotice, callsign, phase)

	if trigger != "" {
		result = join.Empty(
			result,
			fmt.Sprintf(constant.CompactionTrigger, trigger),
		)
	}

	if topic != "" {
		result = join.Empty(
			result,
			fmt.Sprintf(
				constant.CompactionTopic,
				shorten.Prefix(topic, constant.CompactionTopicLimit),
			),
		)
	}

	return result
}
