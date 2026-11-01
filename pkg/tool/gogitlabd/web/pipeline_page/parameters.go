package pipeline_page

func (p *Page) parameters() string {
	return Query(p.view.Instance, p.view.Project, p.view.Pipeline)
}
