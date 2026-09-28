package digest

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func Authorization(
	challenge string,
	method string,
	path string,
	user string,
	password string,
) string {
	realm := challengeField(challenge, "realm")
	nonce := challengeField(challenge, "nonce")
	seed := make([]byte, 8)
	_, _ = rand.Read(seed)
	client := hex.EncodeToString(seed)
	count := "00000001"
	first := hexDigest(user, realm, password)
	second := hexDigest(method, path)
	response := hexDigest(first, nonce, count, client, "auth", second)

	return fmt.Sprintf(
		`Digest username="%s", realm="%s", nonce="%s", uri="%s", qop=auth, nc=%s, cnonce="%s", response="%s", algorithm=MD5`,
		user,
		realm,
		nonce,
		path,
		count,
		client,
		response,
	)
}
