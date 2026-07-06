package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/custom/certificate"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (c *Client) Certificates() []certificate.Certificate {
	list, e := c.dynamic.Resource(constant.CertificateGVR).List(
		c.context,
		v1.ListOptions{},
	)
	errors.PanicOnError(e)
	var result []certificate.Certificate

	for _, item := range list.Items {
		r := certificate.Certificate{
			Name:      item.GetName(),
			Namespace: item.GetNamespace(),
		}
		conditions, _, f := unstructured.NestedSlice(
			item.Object,
			constant.StatusField,
			constant.ConditionsField,
		)
		errors.PanicOnError(f)

		for _, raw := range conditions {
			o, isMap := raw.(map[string]any)

			if !isMap {
				continue
			}

			if o[constant.ConditionFieldType] == constant.ConditionReady &&
				o[constant.ConditionFieldStatus] == constant.ConditionStatusTrue {
				r.Ready = true
			}
		}

		result = append(result, r)
	}

	return result
}
