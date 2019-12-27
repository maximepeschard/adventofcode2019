package image

import "fmt"

// Decode returns a layer representing a decoded image
// using black (0), white (1) and transparent (2) pixels.
func Decode(img *Image) Layer {
	decoded := make(Layer, img.Height)
	for h := 0; h < img.Height; h++ {
		decoded[h] = make([]int, img.Width)
		for w := 0; w < img.Width; w++ {
			decoded[h][w] = -1
		}
	}

	for _, l := range img.Layers {
		for h := 0; h < img.Height; h++ {
			for w := 0; w < img.Width; w++ {
				if decoded[h][w] != 0 && decoded[h][w] != 1 {
					decoded[h][w] = l[h][w]
				}
			}
		}
	}

	return decoded
}

// PrintDecoded displays a layer by printing 1s as Xs
// and everything else as whitespaces.
func PrintDecoded(l Layer) {
	for _, row := range l {
		for _, digit := range row {
			if digit == 1 {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
