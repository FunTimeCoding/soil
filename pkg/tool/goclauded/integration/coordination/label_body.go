package coordination

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

func labelBody(
	key string,
	value string,
	from string,
) client.LabelRequest {
	return client.LabelRequest{Key: key, Value: &value, From: &from}
}
