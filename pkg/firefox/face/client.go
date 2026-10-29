package face

import (
	"github.com/funtimecoding/soil/pkg/firefox/capture"
	"github.com/funtimecoding/soil/pkg/firefox/content"
	"github.com/funtimecoding/soil/pkg/firefox/download"
	"github.com/funtimecoding/soil/pkg/firefox/element"
	"github.com/funtimecoding/soil/pkg/firefox/evaluation"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
	"net/http"
)

type Client interface {
	http.Handler
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
	WaitFor(
		identifier int,
		selector string,
		timeout int,
	) (*tab.Tab, error)
	Evaluate(
		identifier int,
		expression string,
	) (*evaluation.Evaluation, error)
	Click(
		identifier int,
		selector string,
	) (*element.Element, error)
	Fill(
		identifier int,
		selector string,
		value string,
		direct bool,
	) (*element.Element, error)
	Screenshot(identifier int) (*capture.Capture, error)
	DownloadFile(
		l string,
		path string,
	) (*download.Download, error)
	DownloadStatus(identifier int) ([]*download.Download, error)
	Close()
}
