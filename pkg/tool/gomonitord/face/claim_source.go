package face

import "github.com/funtimecoding/soil/pkg/tool/gomonitord/store/claim"

type ClaimSource interface {
	Claims() ([]*claim.Claim, error)
}
