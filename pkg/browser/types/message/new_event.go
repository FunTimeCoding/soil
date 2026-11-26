package message

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/browser/constant"
)

func NewEvent(
	method string,
	parameters string,
) *Reply {
	return &Reply{
		Type:       constant.EventType,
		Method:     method,
		Parameters: json.RawMessage(parameters),
	}
}
