package bluetooth

import "encoding/json"

type section struct {
	Connected    []map[string]json.RawMessage `json:"device_connected"`
	Disconnected []map[string]json.RawMessage `json:"device_not_connected"`
}
