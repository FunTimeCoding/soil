package mock_protocol

import "github.com/funtimecoding/soil/pkg/notation"

func (p *Protocol) Evaluate(
	_ string,
	out any,
) error {
	if p.evaluation == nil {
		return nil
	}

	return notation.DecodeBytes(notation.Marshal(p.evaluation), out)
}
