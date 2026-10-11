package model_context

import (
	"fmt"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
)

func messageHeader(m *message.Message) string {
	to := m.ToName

	if to == "" {
		to = constant.MessageEveryone
	}

	return fmt.Sprintf(
		constant.MessageHeader,
		m.Identifier,
		m.FromName,
		to,
		m.CreatedAt.Local().Format(timeConstant.DateMinute),
	)
}
