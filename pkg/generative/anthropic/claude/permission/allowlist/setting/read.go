package setting

import (
	"encoding/json"
	"os"
)

func Read(path string) (*File, error) {
	b, e := os.ReadFile(path)

	if e != nil {
		return nil, e
	}

	var result File

	if f := json.Unmarshal(b, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
