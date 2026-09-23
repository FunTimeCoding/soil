package web

import "fmt"

func decisionIdentity(identifier uint) string {
	return fmt.Sprintf("decision-%d", identifier)
}
