package parser

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func kind(text string) string {
	switch {
	case strings.HasPrefix(text, constant.TableRowPrefix):
		return constant.BlockTable
	case constant.ListItemPattern.MatchString(text):
		return constant.BlockList
	}

	return constant.BlockParagraph
}
