package fake_extension

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/firefox/types/message"
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
