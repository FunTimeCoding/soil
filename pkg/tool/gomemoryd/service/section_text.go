package service

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func sectionText(
	heading string,
	body []string,
) string {
	return join.Empty(heading, constant.Unix, constant.Unix, join.NewLine(body))
}
