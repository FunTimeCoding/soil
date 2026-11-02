package reach

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/lint/reach/transcript"
)

func blocks(raw []byte) (string, []*transcript.Block) {
	var l transcript.Line

	if json.Unmarshal(raw, &l) != nil || len(l.Message.Content) == 0 {
		return "", nil
	}

	var result []*transcript.Block

	if json.Unmarshal(l.Message.Content, &result) != nil {
		return "", nil
	}

	return l.Cwd, result
}
