package fixture

type Wrapped struct {
	Name  string
	Inner *Wrapped
	Raw   *Primitives
}
