package upstream_tester

import "github.com/funtimecoding/soil/pkg/proxmox/unit/upstream_tester/answer"

func (u *Upstream) AnswerMessage(
	path string,
	line string,
	body string,
) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.answers[path] = answer.New(line, body)
}
