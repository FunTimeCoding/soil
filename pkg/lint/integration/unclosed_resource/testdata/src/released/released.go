package example

import (
	"example/errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sync"
)

type Direct struct {
	resource *Resource
}

func (d *Direct) Close() {
	d.resource.Close()
}

func NewDirect() *Direct {
	return &Direct{resource: open()}
}

type Helped struct {
	resource *Resource
}

func (h *Helped) Close() {
	h.shutdown()
}

func (h *Helped) shutdown() {
	h.resource.Close()
}

func NewHelped() *Helped {
	return &Helped{resource: open()}
}

type Service struct {
	resource *Resource
}

func (s *Service) Close() {
	s.resource.Close()
}

func (s *Service) Stop() {
	s.Close()
}

func newService() *Service {
	return &Service{resource: open()}
}

func UseService() {
	s := newService()
	defer s.Stop()
}

type Owner struct {
	service *Service
}

func (o *Owner) Close() {
	o.service.Stop()
}

func NewOwner() *Owner {
	return &Owner{service: newService()}
}

type Stream struct{}

func (s *Stream) Close() error {
	return nil
}

func openStream() *Stream {
	return &Stream{}
}

type Closer struct {
	stream *Stream
}

func (c *Closer) Close() {
	errors.LogClose(c.stream)
}

func NewCloser() *Closer {
	return &Closer{stream: openStream()}
}

type Bag struct {
	resource *Resource
}

func Wrap(r *Resource) *Bag {
	return &Bag{resource: r}
}

type Source struct {
	resource *Resource
}

func (s *Source) Resource() *Resource {
	return s.resource
}

func Derive(s *Source) *Bag {
	return &Bag{resource: s.Resource()}
}

type Registrar struct{}

func (r *Registrar) Cleanup(func()) {}

func Tester(t *Registrar) *Bag {
	r := open()
	t.Cleanup(r.Close)

	return &Bag{resource: r}
}

func arranged(t *Registrar) *Resource {
	r := open()
	t.Cleanup(r.Close)

	return r
}

func UseArranged(t *Registrar) *Bag {
	return &Bag{resource: arranged(t)}
}

type Pool struct {
	members map[string]*Resource
}

func (p *Pool) Close() {
	for _, m := range p.members {
		m.Close()
	}
}

func (p *Pool) Add(key string) {
	p.members[key] = open()
}

type Alert struct{}

func (a *Alert) Close(string) {}

type Alerts struct {
	alert *Alert
}

func NewAlerts() *Alerts {
	return &Alerts{alert: &Alert{}}
}

type Base struct {
	resource *Resource
}

func (b *Base) Close() {
	b.resource.Close()
}

func newBase() *Base {
	return &Base{resource: open()}
}

type Outer struct {
	*Base
}

func NewOuter() *Outer {
	b := newBase()

	return &Outer{Base: b}
}

type Cleaned struct {
	resource *Resource
}

func NewCleaned(t *Registrar) *Cleaned {
	result := &Cleaned{resource: open()}
	t.Cleanup(result.resource.Close)

	return result
}

func NewDeferred() int {
	c := &Cleaned{resource: open()}
	defer c.resource.Close()

	return 0
}

func shut(r *Resource) {
	r.Close()
}

func shutDeep(r *Resource) {
	shut(r)
}

type Shut struct {
	resource *Resource
}

func (s *Shut) Close() {
	shutDeep(s.resource)
}

func NewShut() *Shut {
	return &Shut{resource: open()}
}

func UseShut() {
	r := open()
	shutDeep(r)
}

var shared = sync.OnceValue(open)

func UseShared() {
	r := shared()
	_ = r
}

func NewSharedBag() *Bag {
	return &Bag{resource: shared()}
}

func (s *Service) Halt(verbose bool) {
	_ = verbose
	s.Close()
}

func UseHalt() {
	s := newService()
	defer s.Halt(true)
}

func graceful(
	c io.Closer,
	verbose bool,
) {
	_ = verbose
	errors.LogClose(c)
}

type Server struct {
	stream *Stream
}

func (s *Server) Stop(verbose bool) {
	graceful(s.stream, verbose)
}

func (s *Server) Close() {
	s.Stop(false)
}

func NewServer() *Server {
	return &Server{stream: openStream()}
}

func NewDisposedBag() *Bag {
	r := open()
	defer shut(r)

	return &Bag{resource: r}
}

func NewShutInline() int {
	c := &Cleaned{resource: open()}
	defer shut(c.resource)

	return 0
}

func (r *Resource) Name() string {
	return ""
}

type Keeper struct {
	resource *Resource
}

func (k *Keeper) Close() {
	k.resource.Close()
}

func (k *Keeper) Last() *Resource {
	return k.resource
}

func lastOf(k *Keeper) *Resource {
	if r := k.Last(); r != nil {
		return r
	}

	return nil
}

func UseLent(k *Keeper) string {
	r := lastOf(k)

	return r.Name()
}

type request struct{}

func (q *request) Execute() (*http.Response, error) {
	response, e := http.Get("http://example.com")

	if e != nil {
		return response, e
	}

	_ = response.Body.Close()

	return response, nil
}

func UseExecute() int {
	q := &request{}
	r, _ := q.Execute()

	return r.StatusCode
}

func (q *request) Run() (*http.Response, error) {
	return q.Execute()
}

func UseRun() int {
	q := &request{}
	r, _ := q.Run()

	return r.StatusCode
}

type Unit interface {
	Close()
}

type adapter struct {
	face Unit
	once sync.Once
}

func newAdapter(u Unit) *adapter {
	return &adapter{face: u}
}

func (a *adapter) Close() {
	a.once.Do(func() { a.face.Close() })
}

type Runner struct {
	units []Unit
}

func (r *Runner) Close() {
	for i := len(r.units) - 1; i >= 0; i-- {
		r.units[i].Close()
	}
}

func (r *Runner) With(u Unit) *Runner {
	v := newAdapter(u)
	r.units = append(r.units, v)

	return r
}

func (r *Runner) WithDirect(u Unit) *Runner {
	r.units = append(r.units, u)

	return r
}

func UseRunner() {
	r := &Runner{}
	defer r.Close()
	s := newService()
	r.With(s)
	d := open()
	r.WithDirect(d)
}

type Keeper2 struct {
	resource *Resource
}

func (k *Keeper2) Close() {
	k.resource.Close()
}

func (k *Keeper2) Resource() *Resource {
	return k.resource
}

func UseGetter(k *Keeper2) string {
	r := k.Resource()

	return r.Name()
}

type Cache struct {
	resources map[string]*Resource
}

func (c *Cache) Close() {
	for _, r := range c.resources {
		r.Close()
	}
}

func (c *Cache) Get(name string) *Resource {
	if r, okay := c.resources[name]; okay {
		return r
	}

	r := open()
	c.resources[name] = r

	return r
}

func UseCache(c *Cache) string {
	r := c.Get("alfa")

	return r.Name()
}

func (c *Cache) Lookup(name string) (*Resource, error) {
	var zero *Resource

	if name == "" {
		return zero, nil
	}

	return c.Get(name), nil
}

type Front struct {
	cache *Cache
}

func (f *Front) Resource(name string) (*Resource, error) {
	return f.cache.Lookup(name)
}

func UseFront(f *Front) string {
	r, _ := f.Resource("alfa")

	return r.Name()
}

type Generic[C Unit] struct {
	clients map[string]C
	build   func() C
}

func (g *Generic[C]) Close() {
	for _, c := range g.clients {
		c.Close()
	}
}

func (g *Generic[C]) Client(name string) (C, error) {
	if c, okay := g.clients[name]; okay {
		return c, nil
	}

	result := g.build()
	g.clients[name] = result

	return result, nil
}

func UseGeneric(g *Generic[*Resource]) string {
	r, _ := g.Client("alfa")

	return r.Name()
}

type Chain struct {
	resource *Resource
}

func newChain() *Chain {
	return &Chain{resource: open()}
}

func (c *Chain) Close() {
	c.resource.Close()
}

func (c *Chain) With() *Chain {
	return c
}

func (c *Chain) Run() {
	c.Close()
}

func UseChain() {
	b := newChain()
	b.With().Run()
}

func UseChainBuilt() {
	b := newChain().With()
	b.Run()
}

func UseInvoked() {
	r := open()
	go func(x *Resource) {
		defer x.Close()
	}(r)
}

type Boxed struct {
	value any
}

func (b *Boxed) Resource() *Resource {
	return b.value.(*Resource)
}

func UseBoxed(b *Boxed) string {
	r := b.Resource()

	return r.Name()
}

type Opener interface {
	Open() *Resource
}

func UseOpener(o Opener) string {
	r := o.Open()

	return r.Name()
}

func Feed() *exec.Cmd {
	f, _ := os.Open("example")

	return &exec.Cmd{Stdin: f}
}
