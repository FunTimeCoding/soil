package guard

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func readInput() *input {
	var result input
	errors.PanicOnError(json.NewDecoder(os.Stdin).Decode(&result))

	return &result
}
