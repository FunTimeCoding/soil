package ladybird

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/types/response"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func scriptResult(
	raw json.RawMessage,
	e error,
) (json.RawMessage, error) {
	v, f := decodeResult[response.Evaluation](raw, e)

	if f != nil {
		return nil, f
	}

	if v.Type == constant.ExceptionResult {
		return nil, validation.New("script exception: %s", v.Exception.Text)
	}

	return v.Result, nil
}
