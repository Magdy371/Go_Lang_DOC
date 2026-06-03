package main

import "fmt"

func slicesToObjects(colorNames []string, hexValues []int) []Color {
	color2 := make([]Color, 0, len(colorNames))

	for i := 0; i < len(colorNames); i++ {
		color2 = append(color2, Color{
			Name: colorNames[i],
			Hex:  hexValues[i],
		})
	}
	return color2
}

type Color struct {
	Name string
	Hex  int
}

func main() {
	var colorName []string = []string{"red", "blue", "green"}
	var hexValues []int = []int{0xFF0000, 0x00FF00, 0x0000FF}
	var colors []Color = slicesToObjects(colorName, hexValues)
	fmt.Println(colors)
}
