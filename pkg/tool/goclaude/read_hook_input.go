package goclaude

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func readHookInput() *HookInput {
	var input HookInput
	errors.PanicOnError(json.NewDecoder(os.Stdin).Decode(&input))

	return &input
}
