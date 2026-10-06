package convert

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func Memories(memories []record.Memory) []*SlimMemory {
	result := make([]*SlimMemory, 0, len(memories))

	for i := range memories {
		result = append(result, Memory(&memories[i]))
	}

	return result
}
