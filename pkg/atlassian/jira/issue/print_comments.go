package issue

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/text"
	"github.com/funtimecoding/soil/pkg/text/option"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"slices"
)

func (i *Issue) PrintComments() {
	if i.Raw.Fields == nil {
		return
	}

	if i.Raw.Fields.Comments == nil {
		return
	}

	for _, c := range i.Raw.Fields.Comments.Comments {
		if slices.Contains(i.commentNameFilter, c.Author.Name) {
			continue
		}

		console.Format(
			"  Comment: %s | %s | %s\n",
			c.Author.Name,
			CommentTime(c).Format(constant.DateMinute),
			consoleConstant.Magenta(
				"%s",
				text.OptimizeWhitespace(c.Body, option.Compact()),
			),
		)
	}
}
