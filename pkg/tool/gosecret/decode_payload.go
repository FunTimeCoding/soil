package gosecret

import (
	"encoding/base64"
	"fmt"
)

func decodePayload(payload map[string]string) (map[string]string, error) {
	result := map[string]string{}

	for k, v := range payload {
		b, e := base64.StdEncoding.DecodeString(v)

		if e != nil {
			return nil, fmt.Errorf("decode key %s: %w", k, e)
		}

		result[k] = string(b)
	}

	return result, nil
}
