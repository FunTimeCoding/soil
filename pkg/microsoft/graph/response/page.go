package response

import "encoding/json"

type Page struct {
	Value    []json.RawMessage `json:"value"`
	NextLink string            `json:"@odata.nextLink"`
}
