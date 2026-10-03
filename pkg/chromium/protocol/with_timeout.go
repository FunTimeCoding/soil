package protocol

import "time"

func (p *Protocol) WithTimeout(d time.Duration) *Protocol {
	p.timeout = d

	return p
}
