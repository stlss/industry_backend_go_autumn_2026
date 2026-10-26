package main

func rotateRunes(s string, shift int) string {
	if s == "" {
		return ""
	}

	initRunes := []rune(s)
	n := len(initRunes)

	shift %= n

	if shift < 0 {
		shift += n
	}

	rotatedRunes := append(initRunes[shift:], initRunes[:shift]...)
	return string(rotatedRunes)
}
