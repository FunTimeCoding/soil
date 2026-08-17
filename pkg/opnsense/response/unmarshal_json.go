package response

import "bytes"

func (f *Flag) UnmarshalJSON(b []byte) error {
	v := string(bytes.Trim(b, `"`))
	*f = v == "1" || v == "true"

	return nil
}
