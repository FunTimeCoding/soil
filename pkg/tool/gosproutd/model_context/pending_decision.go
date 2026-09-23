package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func pendingDecision(d *decision.Decision) map[string]any {
	result := map[string]any{
		"identifier": d.Identifier,
		"question":   d.Question,
		"state":      d.State,
	}

	if d.Answer != "" {
		result[constant.AnswerField] = d.Answer
		result["answer_kind"] = d.AnswerKind
	}

	var replies []string

	for _, v := range d.Turns {
		if v.Author == constant.AuthorUser {
			replies = append(replies, v.Content)
		}
	}

	if replies != nil {
		result["replies"] = replies
	}

	return result
}
