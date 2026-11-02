package upstream_tester

import "github.com/funtimecoding/soil/pkg/strings/join"

func (u *Upstream) AnswerFor(
	method string,
	path string,
	line string,
	body string,
) {
	u.AnswerWith(join.Space(method, path), line, body)
}
