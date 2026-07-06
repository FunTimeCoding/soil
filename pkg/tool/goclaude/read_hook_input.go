package goclaude

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func readHookInput() *hookInput {
	var input hookInput
	errors.PanicOnError(json.NewDecoder(os.Stdin).Decode(&input))

	return &input
}
