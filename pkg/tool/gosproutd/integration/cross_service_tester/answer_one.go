package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
)

func (o *Tester) AnswerOne(session string) {
	o.t.Helper()
	d, e := o.Service.PushDecision(
		session,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	assert.FatalOnError(o.t, e)
	o.Service.Answer(d.Identifier, constant.AnswerKindChoice, "narrow")
}
