package unit

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/testing"
)

func refusingClientset(refusal error) *client.Client {
	k := fake.NewClientset()
	k.PrependReactor(
		"*",
		"*",
		func(testing.Action) (bool, runtime.Object, error) {
			return true, nil, refusal
		},
	)

	return client.FromInterface(k, constant.InCluster)
}
