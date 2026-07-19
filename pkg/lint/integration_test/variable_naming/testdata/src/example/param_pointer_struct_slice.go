package example

type Searched struct{}

func ParamPointerStructSlice(v []*Searched) {
	_ = v
}
