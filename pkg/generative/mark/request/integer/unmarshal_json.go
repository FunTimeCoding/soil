package integer

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"strconv"
)

func (i *Integer) UnmarshalJSON(b []byte) error {
	var n float64

	if e := json.Unmarshal(b, &n); e == nil {
		*i = Integer(int(n))

		return nil
	}

	var s string

	if e := json.Unmarshal(b, &s); e == nil {
		v, f := strconv.Atoi(s)

		if f != nil {
			return validation.New("cannot parse %q as integer", s)
		}

		*i = Integer(v)

		return nil
	}

	return validation.New("cannot unmarshal %s as integer", string(b))
}
