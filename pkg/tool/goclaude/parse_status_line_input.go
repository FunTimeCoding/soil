package goclaude

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
)

func parseStatusLineInput(body []byte) *statusLineInput {
	var input statusLineInput
	errors.PanicOnError(json.Unmarshal(body, &input))

	return &input
}
