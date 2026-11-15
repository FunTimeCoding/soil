package example

import "net/http"

func closeChained(c *Chained) {
	c.Close()
}

type Keeper struct {
	chained *Chained
}

func newKeeper(c *Chained) *Keeper {
	return &Keeper{chained: c}
}

func (k *Keeper) Close() {
	k.chained.Close()
}

type Option struct {
	Chained func() *Chained
}

func newOption(c *Chained) *Option {
	result := &Option{}
	result.Chained = func() *Chained { return c }

	return result
}

func factory(c *Chained) func() *Chained {
	return func() *Chained { return c }
}

func PassedToCloser() {
	closeChained(newChained())
}

func PassedToKeeper() {
	k := newKeeper(newChained())
	defer k.Close()
}

func PassedToOption() *Option {
	return newOption(newChained())
}

func PassedToFactory() func() *Chained {
	return factory(newChained())
}

type Embedder struct {
	*Chained
}

func newEmbedder(c *Chained) *Embedder {
	return &Embedder{Chained: c}
}

func PassedToEmbedder() {
	e := newEmbedder(newChained())
	defer e.Close()
}

type body struct{}

func (b *body) Read([]byte) (int, error) {
	return 0, nil
}

func (b *body) Close() error {
	return nil
}

func newBody() *body {
	return &body{}
}

func Requested() {
	r, _ := http.NewRequest(http.MethodPost, "http://example.com", newBody())
	_ = r
}

func closeLater(c *Chained, register func(func())) {
	register(c.Close)
}

func PassedToRegistrar() {
	closeLater(newChained(), func(func()) {})
}
