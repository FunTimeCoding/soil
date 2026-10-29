package face

import "github.com/funtimecoding/soil/pkg/tool/gomonitord/store/claim"

type Claims interface {
	ClaimSource
	ClaimItem(
		item string,
		owner string,
	) (*claim.Claim, error)
	Release(
		item string,
		owner string,
	) error
}
