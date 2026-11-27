package extension_peer

import "github.com/funtimecoding/soil/pkg/browser/firefox/types/message"

func (x *Extension) Refuse(
	method string,
	text string,
) {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	x.replies[method] = message.NewFailedReply(text)
}
