package example

type Entry struct {
	chained *Chained
}

func (e *Entry) Chained() *Chained {
	return e.chained
}

func (e *Entry) Close() {
	e.chained.Close()
}

type Pool struct {
	entries map[string]*Entry
}

func (p *Pool) Entry(name string) (*Entry, error) {
	return p.entries[name], nil
}

type Pooled struct {
	pool *Pool
}

func (s *Pooled) chained(name string) (*Chained, error) {
	e, f := s.pool.Entry(name)

	if f != nil {
		return nil, f
	}

	return e.Chained(), nil
}

func (s *Pooled) BorrowedFromThePool() error {
	c, e := s.chained("alfa")

	if e != nil {
		return e
	}

	c.Work()

	return nil
}
