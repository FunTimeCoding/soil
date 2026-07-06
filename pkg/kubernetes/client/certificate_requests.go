package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/custom/certificate_request"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (c *Client) CertificateRequests() []certificate_request.CertificateRequest {
	list, e := c.dynamic.Resource(constant.CertificateRequestGVR).List(
		c.context,
		v1.ListOptions{},
	)
	errors.PanicOnError(e)
	var result []certificate_request.CertificateRequest

	for _, item := range list.Items {
		request := certificate_request.CertificateRequest{
			Name:      item.GetName(),
			Namespace: item.GetNamespace(),
			Ready:     true,
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

			if o[constant.ConditionFieldType] != constant.ConditionReady {
				continue
			}

			if o[constant.ConditionFieldStatus] != constant.ConditionStatusTrue {
				request.Ready = false

				if m, isString := o[constant.ConditionFieldMessage].(string); isString {
					request.Message = m
				}
			}
		}

		result = append(result, request)
	}

	return result
}
