package ladybird

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
)

func reply(
	method string,
	r *message.Reply,
) (json.RawMessage, error) {
	if r.Type == constant.LadybirdErrorType {
		return nil, fmt.Errorf("%s: %w", method, refusal(r.Error, r.Message))
	}

	return r.Result, nil
}
