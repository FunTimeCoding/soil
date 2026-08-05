package goclaude

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
)

func parseStatuslineInput(body []byte) *statuslineInput {
	var input statuslineInput
	errors.PanicOnError(json.Unmarshal(body, &input))

	return &input
}
