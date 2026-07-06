package main

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/build"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/markup"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/notifier/mattermost_notifier"
	"github.com/funtimecoding/soil/pkg/project"
	"github.com/funtimecoding/soil/pkg/prometheus/push"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/secure_shell"
	"github.com/funtimecoding/soil/pkg/web"
)

func main() {
	if true {
		return
	}

	build.DebianPackages()
	relational.FirstSkipNotFound(nil, nil, nil)
	unexpected.Float(0)
	markup.ToNotation("")
	system.TransmissionSocket("")
	system.SignalCancelContext(nil)
	system.RunDirectory("", "")
	system.OpenHome("")
	system.ReadLink("")
	system.MustLookup("")
	system.TarWrite(nil, nil)
	system.ReadFull(nil, nil)
	system.WriteFile("", nil, 0)

	if true {
		_, e := system.FindHostnames("")
		errors.PanicOnError(e)
	}

	web.GetBytes(nil, "")
	web.Post(nil, "", "")
	web.Patch(nil, "", "")
	web.PostBytes(nil, "", nil)
	web.Download("", "")
	web.PostFile("", "", "", "")
	web.WriteBytesSafe(nil, 0, nil)
	web.DefaultServer(nil)
	push.Send(nil)
	errors.NotFound("")
	relational.NewEnvironment()
	gitlab.NewGitLabCom("")
	git.ModifiedFiles("")
	metric.MiddlewareServer("", nil)
	secure_shell.Listen(nil, "")
	client.MustNewContext("")
	client.NewInCluster("")
	chromium.NewCombined("")
	brave.OpenProfileLink("", "")
	page.PrintBody(response.Body{})
	strings.PrintTrim("")
	mattermost_notifier.New(nil, "", "")
	project.TemporaryPath("")
}
