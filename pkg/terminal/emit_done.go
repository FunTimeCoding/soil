package terminal

func (t *Terminal) EmitDone(e error) {
	if e != nil {
		t.Exitln(e)
	}
}
