package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/tool/goatlassian/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	s := instrument.New(constant.Identity, version)
	defer func() { s.Flush(recover()) }()
	c := client.NewEnvironment()
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
		PersistentPostRun: func(
			m *cobra.Command,
			_ []string,
		) {
			s.RecordCommand(m.Name())
		},
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
	o.AddCommand(createIssue(c))
	o.AddCommand(updateIssue(c))
	o.AddCommand(getCreateMeta(c))
	o.AddCommand(searchUsers(c))
	o.AddCommand(linkIssues(c))
	o.AddCommand(deleteLink(c))
	o.AddCommand(getLinkTypes(c))
	o.AddCommand(updateComment(c))
	o.AddCommand(deleteComment(c))
	o.AddCommand(getChecklist(c))
	o.AddCommand(addChecklistItem(c))
	o.AddCommand(editChecklistItem(c))
	o.AddCommand(deleteChecklistItem(c))
	o.AddCommand(toggleChecklistItem(c))
	o.AddCommand(deletePage(c))
	o.AddCommand(editPage(c))
	o.AddCommand(getPageDraft(c))
	o.AddCommand(listPages(c))
	o.AddCommand(setPageStatus(c))
	errors.PanicOnError(o.Execute())
}
