package build

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/build/origin"
)

func ParseOrigin(output string) *origin.Origin {
	var v origin.Document

	if json.Unmarshal([]byte(output), &v) != nil {
		return nil
	}

	return origin.New(v.Origin.Hash, v.Time)
}
