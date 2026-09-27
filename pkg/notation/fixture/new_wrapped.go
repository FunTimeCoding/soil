package fixture

func NewWrapped(
	name string,
	inner *Wrapped,
	raw *Primitives,
) *Wrapped {
	return &Wrapped{Name: name, Inner: inner, Raw: raw}
}
