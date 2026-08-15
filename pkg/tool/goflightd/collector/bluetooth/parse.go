package bluetooth

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
)

func Parse(text string) map[string]string {
	var v report

	if json.Unmarshal([]byte(text), &v) != nil {
		return nil
	}

	result := make(map[string]string)

	for _, s := range v.Sections {
		for _, group := range s.Disconnected {
			for name := range group {
				result[name] = constant.Disconnected
			}
		}

		for _, group := range s.Connected {
			for name := range group {
				result[name] = constant.Connected
			}
		}
	}

	return result
}
