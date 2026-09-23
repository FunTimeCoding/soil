package store_tester

import (
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
)

func (o *Tester) Answered(question string) {
	o.t.Helper()
	d := o.Store.PushDecision(
		stringConstant.LowerAlfa,
		question,
		"narrow",
		nil,
		nil,
	)
	o.Store.Answer(
		d.Identifier,
		constant.AnswerKindChoice,
		constant.AnswerChannelQueue,
		"narrow",
	)
}
