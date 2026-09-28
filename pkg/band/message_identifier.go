package band

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func messageIdentifier() string {
	seed := make([]byte, 16)
	_, _ = rand.Read(seed)
	s := hex.EncodeToString(seed)

	return fmt.Sprintf(
		"uuid:%s-%s-%s-%s-%s",
		s[0:8],
		s[8:12],
		s[12:16],
		s[16:20],
		s[20:32],
	)
}
