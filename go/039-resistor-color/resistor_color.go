package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	colors := Colors()
	i := 0

	for _, v := range colors {
		if v == color {
			return i
		}
		i++
	}
	return -1
}
