package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)
	// colors["red"] = "##ff0000"
	// colors["green"] = "##00ff00"
	// colors["blue"] = "##0000ff"

	// colors := make(map[int]string)
	// colors[7] = "#ffffff"

	// delete(colors, 7)
	// fmt.Println(colors)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
