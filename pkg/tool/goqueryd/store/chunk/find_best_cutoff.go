package chunk

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func findBestCutoff(
	points []breakPoint,
	fences []codeFence,
	target int,
) int {
	window := target - constant.ChunkWindow
	bestScore := -1.0
	bestPosition := target

	for _, b := range points {
		if b.position < window {
			continue
		}

		if b.position > target {
			break
		}

		if insideCodeFence(b.position, fences) {
			continue
		}

		distance := target - b.position
		normalized := float64(distance) / float64(constant.ChunkWindow)
		multiplier := 1.0 - (normalized*normalized)*0.7
		score := float64(b.score) * multiplier

		if score > bestScore {
			bestScore = score
			bestPosition = b.position
		}
	}

	return bestPosition
}
