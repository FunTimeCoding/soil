package bookmark

import "github.com/funtimecoding/soil/pkg/notation"

func Parse(b []byte) (*File, error) {
	result := &File{}

	if e := notation.DecodeBytes(b, result); e != nil {
		return nil, e
	}

	return result, nil
}
