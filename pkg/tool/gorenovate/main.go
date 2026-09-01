package gorenovate

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	kubernetes "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/tool/gorenovate/constant"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse(version, gitHash, buildDate)
	var missing []string

	if c := parseConfiguration(); c != nil {
		missing = missingFiles(c)

		for _, f := range missing {
			console.Format("matchFiles not found: %s\n", f)
		}
	}

	k := client.NewEnvironment()
	f := kubernetes.Format

	for _, j := range k.CronJobs(kubernetes.RenovateNamespace) {
		console.Format("CronJob: %s\n", j.Format(f))
	}

	for _, j := range k.Jobs(kubernetes.RenovateNamespace) {
		console.Format("Job: %s\n", j.Format(f))
	}

	if len(missing) > 0 {
		os.Exit(1)
	}
}
