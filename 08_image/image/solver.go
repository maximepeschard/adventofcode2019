package image

// CountDigits returns digit counts in an image layer, as a map.
// Keys are the distinct digits, values are the associated counts.
func CountDigits(l Layer) map[int]int {
	counts := make(map[int]int)

	for _, row := range l {
		for _, digit := range row {
			counts[digit]++
		}
	}

	return counts
}

// FindLayerWithFewest returns the index of the layer with
// the fewest ocurrences of the given digit.
func FindLayerWithFewest(img *Image, digit int) int {
	index := -1
	min := -1

	for i, l := range img.Layers {
		counts := CountDigits(l)
		if index == -1 || counts[digit] < min {
			index = i
			min = counts[digit]
		}
	}

	return index
}
