package fosite_client

import "golang.org/x/crypto/bcrypt"

func (c *Client) ValidateSecret(secret string) error {
	return bcrypt.CompareHashAndPassword([]byte(c.Row.Secret), []byte(secret))
}
