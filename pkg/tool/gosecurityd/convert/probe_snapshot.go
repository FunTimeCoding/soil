package convert

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func ProbeSnapshot(v *client.SystemSnapshot) (*server.SystemSnapshot, error) {
	j, e := json.Marshal(v)

	if e != nil {
		return nil, e
	}

	result := &server.SystemSnapshot{}

	if f := json.Unmarshal(j, result); f != nil {
		return nil, f
	}

	return result, nil
}
