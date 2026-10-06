package primitive

type Mode string

func (m Mode) Valid() bool {
	return m != ""
}

type Option struct {
	Mode Mode
}
