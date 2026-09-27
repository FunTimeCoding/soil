package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"strings"
)

func parseLabelChange(body string) (string, string, string) {
	arrow := strings.Index(body, constant.ChangeArrow)

	if arrow < 0 {
		return "", "", ""
	}

	left := body[:arrow]
	now := body[arrow+len(constant.ChangeArrow):]
	key := left
	past := ""

	if space := strings.Index(left, " "); space >= 0 {
		key = left[:space]
		past = left[space+1:]
	}

	if past == constant.UnsetMarker {
		past = ""
	}

	if strings.TrimSpace(now) == constant.UnsetMarker {
		now = ""
	}

	return key, past, now
}
