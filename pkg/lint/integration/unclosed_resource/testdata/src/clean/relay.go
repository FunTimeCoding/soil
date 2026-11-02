package example

type Settings struct {
	Make func() *Chained
}

func newSettings(c *Chained) *Settings {
	result := &Settings{}
	result.Make = func() *Chained { return c }

	return result
}

type Service struct {
	chained *Chained
}

func newService(s *Settings) *Service {
	c := s.Make()

	return &Service{chained: c}
}

func (v *Service) Close() {
	v.chained.Close()
}

func newServiceFrom(s *Settings) *Service {
	return newService(s)
}

func relayedHelper(c *Chained) *Service {
	return newServiceFrom(newSettings(c))
}

func Relayed() {
	v := newServiceFrom(newSettings(newChained()))
	defer v.Close()
}

func RelayedThroughHelper() {
	v := relayedHelper(newChained())
	defer v.Close()
}

type Inline struct {
	chained *Chained
}

func newInline(s *Settings) *Inline {
	return &Inline{chained: s.Make()}
}

func (i *Inline) Close() {
	i.chained.Close()
}

func RelayedInline() {
	i := newInline(newSettings(newChained()))
	defer i.Close()
}

func useSettings(s *Settings) {}

func drawnAndRegistered(s *Settings, cleanup func(func())) {
	c := s.Make()
	cleanup(c.Close)
	useSettings(s)
}

func RelayedToRegistrar() {
	drawnAndRegistered(newSettings(newChained()), func(func()) {})
}

type Closer interface {
	Close()
}

type Asserting struct {
	held any
}

func newAsserting(s *Settings) *Asserting {
	c := s.Make()

	return &Asserting{held: c}
}

func (a *Asserting) Close() {
	if c, okay := a.held.(Closer); okay {
		c.Close()
	}
}

func RelayedToAsserting() {
	a := newAsserting(newSettings(newChained()))
	defer a.Close()
}

type Assigned struct {
	chained *Chained
}

func newAssigned(s *Settings) *Assigned {
	result := &Assigned{}
	result.chained = s.Make()

	return result
}

func (a *Assigned) Close() {
	a.chained.Close()
}

func RelayedByAssignment() {
	a := newAssigned(newSettings(newChained()))
	defer a.Close()
}

func serviceWith(c *Chained) *Service {
	s := &Settings{}
	s.Make = func() *Chained { return c }

	return newService(s)
}

func RelayedThroughLocalSettings() {
	v := serviceWith(newChained())
	defer v.Close()
}
