package fake_extension

import "github.com/funtimecoding/soil/pkg/firefox/types/message"

func (x *Extension) Refuse(
	method string,
	text string,
) {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	x.replies[method] = message.NewFailedReply(text)
}
