package finding_group

func New(vulnerability string) *Group {
	return &Group{Vulnerability: vulnerability}
}
