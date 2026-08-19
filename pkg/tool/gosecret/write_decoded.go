package gosecret

import (
	"fmt"
	"os"
)

func WriteDecoded(
	path string,
	decoded map[string]string,
) error {
	if e := os.WriteFile(
		path,
		[]byte(DecodedContent(decoded)),
		0600,
	); e != nil {
		return fmt.Errorf("write file: %w", e)
	}

	return nil
}
