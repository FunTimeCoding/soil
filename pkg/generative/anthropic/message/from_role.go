package message

import (
	"github.com/anthropics/anthropic-sdk-go/v2"
	"github.com/funtimecoding/soil/pkg/errors/panics"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func fromRole(
	role string,
	b anthropic.ContentBlockParamUnion,
) anthropic.MessageParam {
	switch role {
	case constant.RoleUser:
		return anthropic.NewUserMessage(b)
	case constant.RoleAssistant:
		return anthropic.NewAssistantMessage(b)
	}

	panics.Unexpected(role)

	return anthropic.MessageParam{}
}
