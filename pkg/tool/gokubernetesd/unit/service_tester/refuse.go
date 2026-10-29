package service_tester

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/testing"
)

func (t *Tester) refuse(
	verb string,
	e error,
) {
	t.Dynamic.PrependReactor(
		verb,
		"*",
		func(testing.Action) (bool, runtime.Object, error) {
			return true, nil, e
		},
	)
}
