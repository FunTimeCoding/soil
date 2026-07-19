package example

import "encoding/json"

type Alias Named

func EmbeddingWrapper() any {
	aux := &struct{ *Alias }{Alias: (*Alias)(&Named{})}

	if json.Unmarshal([]byte("{}"), aux) != nil {
		return nil
	}

	return aux
}
