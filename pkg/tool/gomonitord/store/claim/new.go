package claim

func New(
	item string,
	owner string,
) *Claim {
	return &Claim{Item: item, Owner: owner}
}
