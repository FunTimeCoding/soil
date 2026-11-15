package directory_tester

import "fmt"

func Distinguished(account string) string {
	return fmt.Sprintf("uid=%s,ou=people,dc=example,dc=test", account)
}
