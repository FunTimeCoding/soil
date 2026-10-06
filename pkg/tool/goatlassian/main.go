package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassian/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	t := terminal.New(s)
	c := client.NewEnvironment()
	o := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	o.AddCommand(searchIssues(c, t))
	o.AddCommand(getIssue(c, t))
	o.AddCommand(listProjects(c, t))
	o.AddCommand(searchPages(c, t))
	o.AddCommand(getPage(c, t))
	o.AddCommand(createPage(c, t))
	o.AddCommand(updatePage(c, t))
	o.AddCommand(listSpaces(c, t))
	o.AddCommand(getPageChildren(c, t))
	o.AddCommand(getTransitions(c, t))
	o.AddCommand(transitionIssue(c, t))
	o.AddCommand(addIssueComment(c, t))
	o.AddCommand(addPageComment(c, t))
	o.AddCommand(createIssue(c, t))
	o.AddCommand(updateIssue(c, t))
	o.AddCommand(getCreateMeta(c, t))
	o.AddCommand(searchUsers(c, t))
	o.AddCommand(linkIssues(c, t))
	o.AddCommand(deleteLink(c, t))
	o.AddCommand(getLinkTypes(c, t))
	o.AddCommand(updateComment(c, t))
	o.AddCommand(deleteComment(c, t))
	o.AddCommand(getChecklist(c, t))
	o.AddCommand(addChecklistItem(c, t))
	o.AddCommand(editChecklistItem(c, t))
	o.AddCommand(deleteChecklistItem(c, t))
	o.AddCommand(toggleChecklistItem(c, t))
	o.AddCommand(deletePage(c, t))
	o.AddCommand(editPage(c, t))
	o.AddCommand(getPageDraft(c, t))
	o.AddCommand(listPages(c, t))
	o.AddCommand(setPageStatus(c, t))
	argument.CobraInstrument(o, s)
	argument.CobraStamp(o, constant.Identity)
	errors.PanicOnError(o.Execute())
}
