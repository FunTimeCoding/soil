package response

import "encoding/json"

// Salt's full_return yields a bare bool instead of the return
// object when a minion does not respond.
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
