package protocol

import "context"

func (p *Protocol) target() (context.Context, error) {
	if p.context != nil && p.context.Err() == nil {
		return p.context, nil
	}

	x, e := p.client.TargetContext(p.identifier)

	if e != nil {
		return nil, e
	}

	p.context = x

	return x, nil
}
