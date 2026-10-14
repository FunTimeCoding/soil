package example

import (
	"io"
	"os"
)

type Bag struct {
	resource *Resource
}

func NewBag() *Bag {
	return &Bag{resource: open()}
}

type Pair struct {
	first  *Resource
	second *Resource
}

func (p *Pair) Close() {
	p.first.Close()
}

func NewPair() *Pair {
	return &Pair{first: open(), second: open()}
}

func (p *Pair) Reopen() {
	p.second = open()
}

func NewPositionalPair() *Pair {
	r := open()

	return &Pair{open(), r}
}

type Hider struct {
	reader io.Reader
}

func (h *Hider) Close() {}

func NewHider() *Hider {
	f, _ := os.Open("example")

	return &Hider{reader: f}
}

type Pool struct {
	members map[string]*Resource
}

func (p *Pool) Close() {}

func (p *Pool) Add(key string) {
	p.members[key] = open()
}

type List struct {
	items []*Resource
}

func (l *List) Close() {}

func (l *List) Add() {
	l.items = append(l.items, open())
}

type Base struct {
	resource *Resource
}

func (b *Base) Close() {
	b.resource.Close()
}

type Outer struct {
	*Base
	extra *Resource
}

func NewOuter() *Outer {
	return &Outer{Base: &Base{resource: open()}, extra: open()}
}

func inspect(*Resource) {}

type Watched struct {
	resource *Resource
}

func (w *Watched) Close() {
	inspect(w.resource)
}

func NewWatched() *Watched {
	return &Watched{resource: open()}
}

type Inner struct {
	resource *Resource
}

func (i *Inner) Close() {
	i.resource.Close()
}

func newInner() *Inner {
	return &Inner{resource: open()}
}

func closeResourceOf(i *Inner) {
	i.resource.Close()
}

type Shell struct {
	inner *Inner
}

func (s *Shell) Close() {
	closeResourceOf(s.inner)
}

func NewShell() *Shell {
	return &Shell{inner: newInner()}
}

func (r *Resource) Name() string {
	return ""
}

func fresh(int) *Resource {
	return open()
}

func UseFresh() string {
	r := fresh(1)

	return r.Name()
}

type Unit interface {
	Close()
}

type Plain struct {
	inner Unit
}

func newPlain(u Unit) *Plain {
	return &Plain{inner: u}
}

type Runner struct {
	plains []*Plain
}

func (r *Runner) Close() {
	for _, p := range r.plains {
		_ = p
	}
}

func (r *Runner) WithPlain(u Unit) *Runner {
	r.plains = append(r.plains, newPlain(u))

	return r
}

func UseRunner() {
	r := &Runner{}
	defer r.Close()
	plain := open()
	r.WithPlain(plain)
}

type Shelf struct {
	items []*Resource
}

func (s *Shelf) Close() {}

func (s *Shelf) Put(r *Resource) {
	s.items = append(s.items, r)
}

func UseShelf() {
	s := &Shelf{}
	defer s.Close()
	shelved := open()
	s.Put(shelved)
}
