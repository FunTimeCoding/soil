package goupload

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/build"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	gitlab "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/packages"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goupload/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"os"
	"path/filepath"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	locatorDefault := ""

	if s := os.Getenv(gitlab.InterfaceLocator); s != "" {
		locatorDefault = s
	}

	projectDefault := ""

	if s := os.Getenv(gitlab.ProjectIdentifier); s != "" {
		projectDefault = s
	}

	tagDefault := ""

	if s := os.Getenv(gitlab.CommitTag); s != "" {
		tagDefault = s
	}

	headerDefault := ""

	if s := os.Getenv(gitlab.JobToken); s != "" {
		headerDefault = fmt.Sprintf(
			"%s=%s",
			gitlab.JobTokenHeader,
			s,
		)
	}

	a := argument.NewInstance(constant.Identity)
	a.String(argument.Locator, locatorDefault, "GitLab API base URL")
	a.String(argument.Project, projectDefault, "Project ID to update to")
	a.String(argument.Tag, tagDefault, "Git tag")
	a.String(
		argument.Header,
		headerDefault,
		"Header for authentication in key=value format",
	)
	a.Parse(version, gitHash, buildDate)
	locator := a.Required(argument.Locator)
	fmt.Printf("Locator: %s\n", locator)
	project := a.Required(argument.Project)
	fmt.Printf("Project: %s\n", project)
	tag := a.Required(argument.Tag)
	fmt.Printf("Tag: %s\n", tag)
	headers := build.Headers(a.GetString(argument.Header))
	var runs int

	for _, name := range build.OutputDirectories() {
		for _, systemArchitecture := range build.SystemArchitectures() {
			if p := build.GuessArchivePath(
				name,
				systemArchitecture,
			); p != "" {
				runs++
				file := filepath.Base(p)
				fmt.Printf("Archive: %s\n", file)
				l := packages.UploadLink(locator, project, name, tag, file)
				fmt.Printf("Link: %s\n", l)
				status, body := web.PutFile(
					l,
					headers,
					system.ReadBytesUnsafe(p),
				)

				if status != http.StatusCreated {
					fmt.Printf("Upload failed: %d %s\n", status, body)
					os.Exit(1)
				}
			}
		}
	}

	if runs == 0 {
		fmt.Println("No archive uploaded")
		os.Exit(1)
	}
}
