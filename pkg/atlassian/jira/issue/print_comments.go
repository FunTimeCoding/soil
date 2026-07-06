package issue

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/text"
	"github.com/funtimecoding/soil/pkg/text/option"
	"github.com/funtimecoding/soil/pkg/time"
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

		fmt.Printf(
			"  Comment: %s | %s | %s\n",
			c.Author.Name,
			CommentTime(c).Format(time.DateMinute),
			console.Magenta(
				"%s",
				text.OptimizeWhitespace(c.Body, option.Compact()),
			),
		)
	}
}
