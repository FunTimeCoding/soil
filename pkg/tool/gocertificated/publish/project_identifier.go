package publish

import "context"

func (p *Publisher) projectIdentifier(x context.Context) (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.project != 0 {
		return p.project, nil
	}

	result, e := p.forge.ResolveProject(x, p.projectPath)

	if e != nil {
		return 0, e
	}

	p.project = result

	return result, nil
}
