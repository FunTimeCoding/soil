package boolean

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"strconv"
)

func (b *Boolean) UnmarshalJSON(y []byte) error {
	var v bool

	if e := json.Unmarshal(y, &v); e == nil {
		*b = Boolean(v)

		return nil
	}

	var s string

	if e := json.Unmarshal(y, &s); e == nil {
		a, f := strconv.ParseBool(s)

		if f != nil {
			return validation.New("cannot parse %q as boolean", s)
		}

		*b = Boolean(a)

		return nil
	}

	return validation.New("cannot unmarshal %s as boolean", string(y))
}
