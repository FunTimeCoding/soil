package packages

import (
	"fmt"
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
		fmt.Printf("project: %+v\n", project)
		fmt.Printf("package: %+v\n", p)
		fmt.Printf("file: %+v\n", f)
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
		fmt.Printf("link: %s\n", result)
	}

	return result
}
