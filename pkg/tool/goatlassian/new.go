package goatlassian

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goatlassian/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/face"
	"github.com/spf13/cobra"
)

func New(
	c face.Client,
	t *terminal.Terminal,
) *cobra.Command {
	result := &cobra.Command{
		Use:   constant.Identity.Usage(),
		Short: constant.Identity.Description(),
	}
	result.AddCommand(searchIssues(c, t))
	result.AddCommand(getIssue(c, t))
	result.AddCommand(listProjects(c, t))
	result.AddCommand(searchPages(c, t))
	result.AddCommand(getPage(c, t))
	result.AddCommand(createPage(c, t))
	result.AddCommand(updatePage(c, t))
	result.AddCommand(listSpaces(c, t))
	result.AddCommand(getPageChildren(c, t))
	result.AddCommand(getTransitions(c, t))
	result.AddCommand(transitionIssue(c, t))
	result.AddCommand(addIssueComment(c, t))
	result.AddCommand(addPageComment(c, t))
	result.AddCommand(createIssue(c, t))
	result.AddCommand(updateIssue(c, t))
	result.AddCommand(getCreateMeta(c, t))
	result.AddCommand(searchUsers(c, t))
	result.AddCommand(linkIssues(c, t))
	result.AddCommand(deleteLink(c, t))
	result.AddCommand(getLinkTypes(c, t))
	result.AddCommand(updateComment(c, t))
	result.AddCommand(deleteComment(c, t))
	result.AddCommand(getChecklist(c, t))
	result.AddCommand(addChecklistItem(c, t))
	result.AddCommand(editChecklistItem(c, t))
	result.AddCommand(deleteChecklistItem(c, t))
	result.AddCommand(toggleChecklistItem(c, t))
	result.AddCommand(deletePage(c, t))
	result.AddCommand(editPage(c, t))
	result.AddCommand(getPageDraft(c, t))
	result.AddCommand(listPages(c, t))
	result.AddCommand(setPageStatus(c, t))

	return result
}
