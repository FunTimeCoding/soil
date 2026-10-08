package guard

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func readInput() *Input {
	var result Input
	errors.PanicOnError(json.NewDecoder(os.Stdin).Decode(&result))

	return &result
}
