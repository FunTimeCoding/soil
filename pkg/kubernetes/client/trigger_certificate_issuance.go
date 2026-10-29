package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"time"
)

func (c *Client) TriggerCertificateIssuance(
	namespace string,
	name string,
) error {
	resource := c.dynamic.Resource(constant.CertificateGVR).Namespace(namespace)
	o, e := resource.Get(c.context, name, v1.GetOptions{})

	if e != nil {
		return refusal.Classify(e)
	}

	existing, _, f := unstructured.NestedSlice(
		o.Object,
		constant.StatusField,
		constant.ConditionsField,
	)

	if f != nil {
		return f
	}

	conditions := make([]any, 0, len(existing)+1)

	for _, raw := range existing {
		condition, isMap := raw.(map[string]any)

		if isMap &&
			condition[constant.ConditionFieldType] == constant.ConditionIssuing {
			continue
		}

		conditions = append(conditions, raw)
	}

	conditions = append(
		conditions,
		map[string]any{
			constant.ConditionFieldType:    constant.ConditionIssuing,
			constant.ConditionFieldStatus:  constant.ConditionStatusTrue,
			constant.ConditionFieldReason:  constant.ConditionReasonTriggered,
			constant.ConditionFieldMessage: constant.ConditionMessageTriggered,
			constant.ConditionFieldTransition: time.Now().UTC().Format(
				time.RFC3339,
			),
		},
	)

	if g := unstructured.SetNestedSlice(
		o.Object,
		conditions,
		constant.StatusField,
		constant.ConditionsField,
	); g != nil {
		return g
	}

	_, h := resource.UpdateStatus(c.context, o, v1.UpdateOptions{})

	return refusal.Classify(h)
}
