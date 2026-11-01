package profile

func (r *Result) WithNotice(notice string) *Result {
	r.Notice = notice
	r.Text = Text(r)

	return r
}
