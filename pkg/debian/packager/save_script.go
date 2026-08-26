package packager

import "github.com/funtimecoding/soil/pkg/system"

func (p *Packager) saveScript(
	name string,
	text string,
) {
	path := p.scriptPath(name)
	system.MakeDirectory(p.ConfigurationRoot)
	system.SaveFile(path, text)
	system.Executable(path)
}
