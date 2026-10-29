package service

import "github.com/funtimecoding/soil/pkg/tool/gomonitord/store/claim"

func (s *Service) ClaimItem(
	item string,
	owner string,
) (*claim.Claim, error) {
	if e := validateClaim(item, owner); e != nil {
		return nil, e
	}

	return s.store.ClaimItem(item, owner)
}
