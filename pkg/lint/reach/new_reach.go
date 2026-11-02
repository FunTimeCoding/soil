package reach

func NewReach(
	packages []string,
	documents []string,
) *Reach {
	return &Reach{Packages: packages, Documents: documents}
}
