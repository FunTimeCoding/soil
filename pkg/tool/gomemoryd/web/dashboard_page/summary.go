package dashboard_page

import "fmt"

func (p *Page) Summary() []string {
	return []string{
		fmt.Sprintf("%d active", p.view.Active),
		fmt.Sprintf("%d relations", p.view.Relations),
	}
}
