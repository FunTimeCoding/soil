package message

import "github.com/anthropics/anthropic-sdk-go/v2"

func (m *Message) ToParameter() anthropic.MessageParam {
	return fromRole(
		m.Role,
		anthropic.ContentBlockParamUnion{
			OfText: &anthropic.TextBlockParam{Text: m.Content},
		},
	)
}
