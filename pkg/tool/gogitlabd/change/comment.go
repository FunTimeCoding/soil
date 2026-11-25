package change

import (
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_event"
	"github.com/funtimecoding/soil/pkg/watch/digest"
	"github.com/funtimecoding/soil/pkg/watch/digest/event"
	"time"
)

func Comment(
	n *note.Note,
	writer string,
	at time.Time,
) *request_event.Event {
	author := n.Author

	if writer != "" {
		author = writer
	}

	if n.Create != nil {
		at = *n.Create
	}

	return request_event.New(
		event.New(
			constant.CommentNoun,
			author,
			join.Empty(author, ": ", digest.Excerpt(n.Body)),
			at,
		),
		writer,
	)
}
