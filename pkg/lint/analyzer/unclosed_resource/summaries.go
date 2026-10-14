package unclosed_resource

type Summaries struct {
	arranges map[string]bool
	disposes map[string]map[int]bool
	releases map[string]map[string]bool
	shared   map[string]bool
	lends    map[string]bool
	self     map[string]bool
	spent    map[string]bool
	wraps    map[string]map[int]bool
}
