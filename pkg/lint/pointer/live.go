package pointer

func (r *Resolution) Live() bool {
	return r.Concern == "" && r.Reason == ""
}
