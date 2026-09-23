//go:build ci

package client

import "fmt"

func distinguished(account string) string {
	return fmt.Sprintf("uid=%s,ou=people,dc=example,dc=test", account)
}
