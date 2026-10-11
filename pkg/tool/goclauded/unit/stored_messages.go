package unit

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"

func storedMessages(
	count int,
	from string,
	body string,
) []*message.Message {
	var result []*message.Message

	for i := 1; i <= count; i++ {
		result = append(result, storedMessage(uint(i), from, body))
	}

	return result
}
