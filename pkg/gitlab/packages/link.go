package packages

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func Link(
	host string,
	project *project.Project,
	p *gitlab.Package,
	f *gitlab.PackageFile,
	verbose bool,
) string {
	if verbose {
		console.Format("project: %+v\n", project)
		console.Format("package: %+v\n", p)
		console.Format("file: %+v\n", f)
	}

	result := locator.New(host).Base(constant.Base).Path(
		"projects/%d/packages/%s/%s/%s/%s",
		project.Identifier,
		p.PackageType,
		p.Name,
		p.Version,
		f.FileName,
	).String()

	if verbose {
		console.Format("link: %s\n", result)
	}

	return result
}
