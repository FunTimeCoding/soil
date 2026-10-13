package evaluation

import "encoding/json"

type Evaluation struct {
	Identifier int             `json:"tab_id"`
	Value      json.RawMessage `json:"value"`
}
