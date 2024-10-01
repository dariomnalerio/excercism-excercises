package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendation := '❗'
	search := '🔍'
	weather := '☀'
	app := "default"

	for _, ch := range log {
		switch ch {
		case recommendation:
			app = "recommendation"
		case search:
			app = "search"
		case weather:
			app = "weather"
		}
		if app != "default" {
			break
		}
	}
	return app
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for i, ch := range runes {
		if ch == oldRune {
			runes[i] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
