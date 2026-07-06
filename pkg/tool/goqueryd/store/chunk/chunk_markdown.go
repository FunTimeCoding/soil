package chunk

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func chunkMarkdown(content string) []Chunk {
	points := scanBreakPoints(content)
	fences := findCodeFences(content)
	position := 0
	var result []Chunk

	for position < len(content) {
		target := position + constant.ChunkSize

		if target > len(content) {
			target = len(content)
		}

		end := target

		if end < len(content) {
			best := findBestCutoff(points, fences, target)

			if best > position && best <= target {
				end = best
			}
		}

		if end <= position {
			end = target
		}

		result = append(
			result,
			Chunk{
				Text:     content[position:end],
				Position: position,
			},
		)

		if end >= len(content) {
			break
		}

		next := end - constant.ChunkOverlap

		if len(result) > 0 && next <= result[len(result)-1].Position {
			next = end
		}

		position = next
	}

	return result
}
