package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/custom/certificate"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/custom/certificate_request"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/cron_job"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/job"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/namespace"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/node"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/pod"
	"io"
	batch "k8s.io/api/batch/v1"
	networking "k8s.io/api/networking/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"time"
)

type Client interface {
	CertificateRequests() (
		[]certificate_request.CertificateRequest,
		error,
	)
	Certificates() ([]certificate.Certificate, error)
	Cluster() string
	Configuration() *rest.Config
	ConfigurationContexts(filterExtended bool) []string
	CreateJobFromCron(
		namespace string,
		cron string,
		name string,
	) (*batch.Job, error)
	CronJob(
		namespace string,
		name string,
	) (*cron_job.Job, error)
	CronJobs(namespace string) ([]*cron_job.Job, error)
	DeleteCertificateRequest(
		namespace string,
		name string,
	) error
	DeleteEvent(f *filter.Filter) error
	DeleteEventSimple(v *event.Event) error
	DeleteJob(
		namespace string,
		name string,
	) error
	DeleteJobWatch(
		namespace string,
		name string,
	) error
	Dynamic() dynamic.Interface
	Events(
		f *filter.Filter,
		limit int,
		fieldSelector string,
	) ([]*event.Event, error)
	EventsSimple(
		normal bool,
		warning bool,
	) ([]*event.Event, error)
	Execute(
		stdout io.Writer,
		stderr io.Writer,
		namespace string,
		pod string,
		container string,
		command ...string,
	) error
	ExecuteSimple(
		namespace string,
		pod string,
		command ...string,
	) (string, string, error)
	ExecuteStandard(
		namespace string,
		pod string,
		container string,
		command ...string,
	) error
	Job(
		namespace string,
		name string,
	) (*job.Job, error)
	JobRunning(
		namespace string,
		name string,
	) (bool, error)
	Jobs(namespace string) ([]*job.Job, error)
	Lease(
		namespace string,
		name string,
	) (*lease.Lease, error)
	Log(
		namespace string,
		pod string,
		container string,
	) (string, error)
	MustCreateJobFromCron(
		namespace string,
		cron string,
		name string,
	) *batch.Job
	MustCronJobs(namespace string) []*cron_job.Job
	MustDeleteEventSimple(v *event.Event)
	MustDeleteJobWatch(
		namespace string,
		name string,
	)
	MustEventsSimple(
		normal bool,
		warning bool,
	) []*event.Event
	MustJob(
		namespace string,
		name string,
	) *job.Job
	MustJobs(namespace string) []*job.Job
	MustLease(
		namespace string,
		name string,
	) *lease.Lease
	MustLog(
		namespace string,
		pod string,
		container string,
	) string
	MustNamespaces(f *filter.Filter) []*namespace.Namespace
	MustNodes() []*node.Node
	MustPods(f *filter.Filter) []*pod.Pod
	Namespaces(f *filter.Filter) ([]*namespace.Namespace, error)
	Nested() kubernetes.Interface
	NetworkPolicy(
		namespace string,
		name string,
	) (*networking.NetworkPolicy, error)
	Nodes() ([]*node.Node, error)
	Pod(
		namespace string,
		name string,
	) (*pod.Pod, error)
	Pods(f *filter.Filter) ([]*pod.Pod, error)
	Reachable(x context.Context) error
	TriggerCertificateIssuance(
		namespace string,
		name string,
	) error
	UpdateNetworkPolicy(np *networking.NetworkPolicy) error
	Validate() error
	WaitForJob(
		namespace string,
		name string,
		limit time.Duration,
	) error
}
