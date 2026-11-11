package example

import "example/codec"

type Decoded struct{}

func (d *Decoded) UnmarshalYAML(node *codec.Node) error {
	return nil
}

var _ codec.Unmarshaler = &Decoded{}
