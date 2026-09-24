package attribution

func NewLoad(
	test string,
	total int,
	alone int,
) *Load {
	return &Load{Test: test, Total: total, Alone: alone}
}
