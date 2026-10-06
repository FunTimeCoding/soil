package example

import "encoding/json"

type Decoded struct{}

func (d *Decoded) UnmarshalJSON(b []byte) error {
	return nil
}

var _ json.Unmarshaler = &Decoded{}
