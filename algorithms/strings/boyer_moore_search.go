package strings



func Contains(text, pattern string) bool {
	return Index(text, pattern) != -1
}

func Index(text, pattern string) int {
	textLen := len(text)
	patternLen := len(pattern)

	if textLen < patternLen {
		return -1
	}

	if textLen == patternLen {
		if text[0] == pattern[0] {
			return 0
		}
		return -1
	}

	return boyerMooreSearch(text, pattern, 0)
}

func boyerMooreSearch(text, pattern string, start int) int {
	curText := text[start:]
	bcPos := findBadCharacterPos(curText, pattern)
	if bcPos == -1 {
		return start
	}
	lcPos := lastCommonPos(pattern, curText[bcPos])
	next := start + bcPos - lcPos
	return boyerMooreSearch(text, pattern, next)
}

func findBadCharacterPos(text, pattern string) int {
	end := len(pattern) - 1

	for text[end] == pattern[end] {
		if end == 0 {
			return -1
		}
		end--
	}
	return end

}

func lastCommonPos(pattern string, badCharacter byte) int {

	for i := 0; i < len(pattern)-1; i++ {
		if pattern[i] == badCharacter {
			return i
		}
	}
	return -1
}
