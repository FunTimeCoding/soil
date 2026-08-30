package notation

import "encoding/json/v2"

func DecodeBytes(
	value []byte,
	structure any,
) error {
	return json.Unmarshal(value, structure)
}
