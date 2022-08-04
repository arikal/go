package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
		"black": "000000",
		"white": "ffffff",
	}

	printColors(colors)
}

func printColors(colors map[string]string) {
	for color, hexCode := range colors {
		fmt.Println("Hex code for", color, "is", hexCode)
	}
}
