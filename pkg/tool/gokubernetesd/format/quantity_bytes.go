package format

import "k8s.io/apimachinery/pkg/api/resource"

func QuantityBytes(raw string) int64 {
	if raw == "" {
		return 0
	}

	q, e := resource.ParseQuantity(raw)

	if e != nil {
		return 0
	}

	return q.Value()
}
