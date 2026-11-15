package label_tester

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

func Body(
	key string,
	value string,
	from string,
) client.LabelRequest {
	return client.LabelRequest{Key: key, Value: &value, From: &from}
}
