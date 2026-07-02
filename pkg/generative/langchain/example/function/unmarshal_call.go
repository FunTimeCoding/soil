package function

import "github.com/funtimecoding/go-library/pkg/notation"

func unmarshalCall(input string) *Call {
	var c Call

	if e := notation.Decode(input, &c); e == nil && c.Tool != "" {
		return &c
	}

	return nil
}
