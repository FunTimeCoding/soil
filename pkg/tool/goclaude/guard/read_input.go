package guard

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/request"
	"os"
)

func readInput() *request.Input {
	var result request.Input
	errors.PanicOnError(json.NewDecoder(os.Stdin).Decode(&result))

	return &result
}
