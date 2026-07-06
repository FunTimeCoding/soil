package download

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/packages"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/common"
	"github.com/funtimecoding/soil/pkg/tool/godownload/download/option"
	library "gitlab.com/gitlab-org/api/client-go/v2"
	"os"
)

func Run(o *option.Download) {
	validate(o)
	g := gitlab.New(o.Host, o.Token)
	project := common.FindProjectOrExit(g, o.Owner, o.Repository)
	filtered := packages.FilterByName(
		g.MustPackages(project.Identifier, false),
		o.Package,
	)
	var p *library.Package

	if o.PackageVersion == constant.LatestVersion {
		p = packages.FindLatest(filtered)
	} else {
		p = packages.FindVersionOrLatest(filtered, o.PackageVersion)
	}

	if p == nil {
		fmt.Printf("package not found: %s\n", o.Package)
		os.Exit(1)
	}

	if o.PackageVersion != constant.LatestVersion &&
		p.Version != o.PackageVersion {
		errors.Warning(
			"version %s not found, using latest\n",
			o.PackageVersion,
		)
	}

	f := packages.SystemMatchingFile(
		g.MustPackageFiles(project.Identifier, p.ID),
	)

	if f == nil {
		fmt.Printf("architecture and system match not found")
		os.Exit(1)
	}

	if !system.FileExists(o.Package) {
		if !system.FileExists(f.FileName) {
			packages.Download(
				packages.Link(o.Host, project, p, f, o.Verbose),
				o.Token,
				f.FileName,
			)
		}

		system.ExtractZip(f.FileName, system.WorkDirectory())
		system.DeleteFile(f.FileName)
	} else {
		fmt.Printf("file already exists: %s\n", o.Package)
	}

	system.Executable(o.Package)

	if o.Output != DefaultOutput {
		system.Move(o.Package, join.Absolute(o.Output, o.Package))
	}
}
