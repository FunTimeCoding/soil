package upstream_tester

func (u *Upstream) AnswerWith(
	path string,
	line string,
	body string,
) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.answers[path] = &Answer{Line: line, Body: body}
}
