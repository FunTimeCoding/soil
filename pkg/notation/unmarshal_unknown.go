package notation

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/structs/constant"
	"reflect"
)

func UnmarshalUnknown(
	b []byte,
	a any,
	unknownField string,
) error {
	var raw map[string]any

	if e := json.Unmarshal(b, &raw); e != nil {
		return e
	}

	if e := json.Unmarshal(b, a); e != nil {
		return e
	}

	known := make(map[string]bool)
	e := reflect.ValueOf(a).Elem()
	t := e.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		j := f.Tag.Get(constant.NotationKey)

		if j != "" && j != "-" {
			// omitempty tags
			for x := 0; x < len(j); x++ {
				if j[x] == ',' {
					j = j[:x]

					break
				}
			}

			known[j] = true
		} else if j == "" {
			// no tag = lowercase field
			known[f.Name] = true
		}
	}

	unknown := make(map[string]any)

	for k, v := range raw {
		if !known[k] {
			unknown[k] = v
		}
	}

	if v := e.FieldByName(unknownField); v.IsValid() && v.CanSet() {
		v.Set(reflect.ValueOf(unknown))
	}

	return nil
}
