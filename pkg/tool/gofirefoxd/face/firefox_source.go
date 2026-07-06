package face

import (
	"github.com/funtimecoding/soil/pkg/firefox/content"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
)

type FirefoxSource interface {
	Tabs() ([]*tab.Tab, error)
	CreateTab(l string) (*tab.Tab, error)
	ReadTab(
		identifier int,
		raw bool,
	) (*content.Content, error)
	CloseTab(identifier int) error
	Navigate(
		identifier int,
		l string,
	) error
	NavigateBack(identifier int) error
	GroupTabs(
		tabIdentifiers []int,
		groupIdentifier int,
		title string,
		color string,
	) (int, error)
	UngroupTab(identifier int) error
	UpdateGroup(
		groupIdentifier int,
		title string,
		color string,
		collapsed *bool,
	) error
	MustTabs() []*tab.Tab
	MustCreateTab(l string) *tab.Tab
	MustReadTab(
		identifier int,
		raw bool,
	) *content.Content
	MustGroupTabs(
		tabIdentifiers []int,
		groupIdentifier int,
		title string,
		color string,
	) int
}
