package firefox

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/firefox/types/message"
)

func decodeResult(
	r *message.Reply,
	v any,
) error {
	return json.Unmarshal(r.Result, v)
}
