package extension_peer

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/browser/firefox/types/message"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (x *Extension) Answer(
	method string,
	result any,
) {
	b, e := json.Marshal(result)
	errors.PanicOnError(e)
	x.mutex.Lock()
	defer x.mutex.Unlock()
	x.replies[method] = message.NewResultReply(b)
}
