package operation

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/batch/v1"
)

func CronJob(
	c kubernetes.Interface,
	namespace string,
) v1.CronJobInterface {
	return c.BatchV1().CronJobs(namespace)
}
