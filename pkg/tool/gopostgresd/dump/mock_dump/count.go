package mock_dump

import "context"

func (p *Program) Count(
	_ context.Context,
	_ string,
) (int, error) {
	return p.Tables, nil
}
