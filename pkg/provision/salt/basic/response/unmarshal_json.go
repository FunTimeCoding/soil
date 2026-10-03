package response

import "encoding/json"

func (r *LocalReturn) UnmarshalJSON(b []byte) error {
	var responded bool

	if json.Unmarshal(b, &responded) == nil {
		r.Responded = responded

		return nil
	}

	type plain LocalReturn
	var p plain

	if e := json.Unmarshal(b, &p); e != nil {
		return e
	}

	*r = LocalReturn(p)
	r.Responded = true

	return nil
}
