package image

import "strconv"

// A Layer represents a 2D grid of pixels
type Layer [][]int

// An Image has a size and consists of multiple layers
type Image struct {
	Width  int
	Height int
	Layers []Layer
}

// Parse returns a valid Image from a string given its size.
func Parse(width int, height int, data string) (*Image, error) {
	size := len(data)
	nbLayers := size / (width * height)
	layers := make([]Layer, nbLayers)
	cursor := 0

	for layerIndex := 0; layerIndex < nbLayers; layerIndex++ {
		layer := make(Layer, height)

		for h := 0; h < height; h++ {
			row := make([]int, width)

			for w := 0; w < width; w++ {
				digit, err := strconv.Atoi(data[cursor : cursor+1])
				if err != nil {
					return nil, err
				}
				row[w] = digit
				cursor++
			}

			layer[h] = row
		}

		layers[layerIndex] = layer
	}

	return &Image{Width: width, Height: height, Layers: layers}, nil
}
