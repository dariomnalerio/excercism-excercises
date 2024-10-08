package reverse

func Reverse(input string) string {
	runes := []rune(input)
	j := len(runes) - 1
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[j] = runes[j], runes[i]
		j--
	}
	return string(runes)
}
