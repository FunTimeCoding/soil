package message

import "encoding/json"

func NewResultReply(result json.RawMessage) *Reply {
	return &Reply{Result: result}
}
