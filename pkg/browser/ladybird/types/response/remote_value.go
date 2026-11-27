package response

import "encoding/json"

type RemoteValue struct {
	Type             string          `json:"type"`
	Value            json.RawMessage `json:"value"`
	SharedIdentifier string          `json:"sharedId"`
}
