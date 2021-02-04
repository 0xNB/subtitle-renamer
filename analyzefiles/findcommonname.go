// Package analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func LcSubstr(strinsg []string) string {
	return ""
}

// LcSubStr computes the longest common substring of two strings.
func LcSubStr(s1 string, s2 string) string {
	// create two dim array that holds our longest common substrings for DP
	l1 := len(s1)
	l2 := len(s2)

	lcDP := make([]int, l1*l2)
	var resultLen int
	var lastIndex int

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if i == 0 || j == 0 {
				lcDP[i*l1+j] = 0
			} else if s1[i-1] == s2[j-1] {
				lcDP[i*l1+j] = lcDP[(i-1)*l1+(j-1)] + 1
				resultLen = max(resultLen, lcDP[i*l1+j])
				lastIndex = i
			} else {
				lcDP[i*l1+j] = 0
			}

		}
	}

	// convert to go runes and return the slice of those runes
	runes := []rune(s1)
	return string(runes[lastIndex : lastIndex+resultLen])

}

func deduceCommonName(files []AnalyzedFile) (bool, string) {
	// TOOD impl
	return false, ""
}
