package terminal

func (t *Terminal) Reject(
	status string,
	body []byte,
) {
	if len(body) == 0 {
		t.Exitln(status)

		return
	}

	t.Exitf("%s: %s\n", status, body)
}
