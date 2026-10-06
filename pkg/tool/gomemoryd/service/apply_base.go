package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maps"
)

func applyBase(
	metadata map[string]string,
	base *string,
) map[string]string {
	if base == nil {
		return metadata
	}

	result := maps.Clone(metadata)

	if *base == "" {
		delete(result, constant.BaseKey)

		return result
	}

	if result == nil {
		result = map[string]string{}
	}

	result[constant.BaseKey] = *base

	return result
}
