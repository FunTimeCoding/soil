package convert

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func Memories(memories []store.Memory) []*SlimMemory {
	result := make([]*SlimMemory, 0, len(memories))

	for i := range memories {
		result = append(result, Memory(&memories[i]))
	}

	return result
}
