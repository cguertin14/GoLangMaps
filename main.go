package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string

	// colors := make(map[int]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	// colors[10] = "#ffffff"

	// delete(colors, 10)

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
