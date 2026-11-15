package delivery_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"

func StoredMessages(
	count int,
	from string,
	body string,
) []*message.Message {
	var result []*message.Message

	for i := 1; i <= count; i++ {
		result = append(result, StoredMessage(uint(i), from, body))
	}

	return result
}
