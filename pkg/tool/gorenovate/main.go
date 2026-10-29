package gorenovate

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	kubernetesConstant "github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/tool/gorenovate/constant"
	"os"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse()
	var missing []string

	if c := parseConfiguration(); c != nil {
		missing = missingFiles(c)

		for _, f := range missing {
			console.Format("matchFiles not found: %s\n", f)
		}
	}

	k := client.NewEnvironment()
	f := kubernetesConstant.Format

	for _, j := range k.MustCronJobs(kubernetesConstant.RenovateNamespace) {
		console.Format("CronJob: %s\n", j.Format(f))
	}

	for _, j := range k.MustJobs(kubernetesConstant.RenovateNamespace) {
		console.Format("Job: %s\n", j.Format(f))
	}

	if len(missing) > 0 {
		os.Exit(1)
	}
}
