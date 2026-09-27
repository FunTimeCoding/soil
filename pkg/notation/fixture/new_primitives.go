package fixture

func NewPrimitives(
	s string,
	i int,
	f float64,
	b bool,
) *Primitives {
	return &Primitives{String: s, Integer: i, Float: f, Boolean: b}
}
