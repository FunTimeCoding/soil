package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassian/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	c := client.NewEnvironment()
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(searchIssues(c))
	o.AddCommand(getIssue(c))
	o.AddCommand(listProjects(c))
	o.AddCommand(searchPages(c))
	o.AddCommand(getPage(c))
	o.AddCommand(createPage(c))
	o.AddCommand(updatePage(c))
	o.AddCommand(listSpaces(c))
	o.AddCommand(getPageChildren(c))
	o.AddCommand(getTransitions(c))
	o.AddCommand(transitionIssue(c))
	o.AddCommand(addIssueComment(c))
	o.AddCommand(addPageComment(c))
	errors.PanicOnError(o.Execute())
}
