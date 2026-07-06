package protocol

import "github.com/funtimecoding/soil/pkg/chromium"

func New(tab string) *Protocol {
	c := chromium.NewEnvironment()
	t := c.TabByHost(tab)

	if t == nil {
		panic("tab not found")
	}

	return &Protocol{client: c, context: c.TargetContext(t.Identifier)}
}
