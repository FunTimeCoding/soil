package goclaude

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
)

func parseStatusLineInput(body []byte) *StatusLineInput {
	var input StatusLineInput
	errors.PanicOnError(json.Unmarshal(body, &input))

	return &input
}
