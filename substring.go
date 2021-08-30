package substring

const emptyString = ``

// Substring returns the substring of string s between startIndex (inclusive)
// and endIndex (exclusive). The substring is based on the number of indiviual
// runes in the string; not based on bytes!
// If the startIndex is greater than the number of runes an empty string is
// returned. If startIndex == endIndex or startIndex + 1 == endIndex an empty
// string is returned. If startIndex > endIndex and empty string is returend.
func Substring(s string, startIndex, endIndex int) string {
	runes := []rune(s)
	length := len(runes)

	if startIndex > length {
		return emptyString
	}
	if startIndex == endIndex || startIndex+1 == endIndex {
		return emptyString
	}

	if startIndex < 0 {
		startIndex = 0
	}
	// Check this here, for cases in which the provided indices are both negative.
	// e.g. startIndex = -10, endIndex = -1
	if startIndex > endIndex {
		return emptyString
	}
	if endIndex > length {
		endIndex = length
	}
	return string(runes[startIndex:endIndex])
}

// SubstringEnd returns the substring based on runes until the endIndex (exclusive).
// If the endIndex is smaller than 0 an empty string is returned. If the endIndex is
// larger than the number of runes in the string, the string is returned unmodified.
func SubstringEnd(s string, endIndex int) string {
	return Substring(s, 0, endIndex)
}

// SubstringStart returns the substring based on runes from the startIndex (inclusive)
// until the end of the string. If the startIndex is larger than the number of
// runes in the string an empty string is returned.
func SubstringStart(s string, startIndex int) string {
	return Substring(s, startIndex, len(s))
}
