package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/custom/certificate"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (c *Client) Certificates() ([]certificate.Certificate, error) {
	list, e := c.dynamic.Resource(constant.CertificateGVR).List(
		c.context,
		v1.ListOptions{},
	)

	if e != nil {
		return nil, refusal.Classify(e)
	}

	var result []certificate.Certificate

	for _, item := range list.Items {
		issuer, _, f := unstructured.NestedString(
			item.Object,
			constant.SpecField,
			constant.IssuerReferenceField,
			constant.IssuerReferenceName,
		)

		if f != nil {
			return nil, f
		}

		r := certificate.Certificate{
			Name:      item.GetName(),
			Namespace: item.GetNamespace(),
			Issuer:    issuer,
		}
		conditions, _, g := unstructured.NestedSlice(
			item.Object,
			constant.StatusField,
			constant.ConditionsField,
		)

		if g != nil {
			return nil, g
		}

		for _, raw := range conditions {
			o, isMap := raw.(map[string]any)

			if !isMap {
				continue
			}

			if o[constant.ConditionFieldType] == constant.ConditionReady &&
				o[constant.ConditionFieldStatus] == constant.ConditionStatusTrue {
				r.Ready = true
			}

			if o[constant.ConditionFieldType] == constant.ConditionIssuing {
				r.Issuing = true
			}
		}

		result = append(result, r)
	}

	return result, nil
}
