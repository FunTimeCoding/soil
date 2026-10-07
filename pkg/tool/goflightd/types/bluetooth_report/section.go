package bluetooth_report

import "encoding/json"

type Section struct {
	Connected    []map[string]json.RawMessage `json:"device_connected"`
	Disconnected []map[string]json.RawMessage `json:"device_not_connected"`
}
