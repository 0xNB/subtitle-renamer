// Package analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func LcSubstr(strings []string) string {
	return ""
}

// LcSubStr computes the longest common substring of two strings using dynamic programming
func LcSubStr(s1 string, s2 string) (string, int) {
	
	l1 := len(s1)
	l2 := len(s2)

	lcDP := make([]int, l1*l2)
	var resultLen int
	var lastIndex int

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if i == 0 || j == 0 {
				lcDP[j*l1+i] = 0
			} else if s1[i-1] == s2[j-1] {
				lcDP[j*l1+i] = lcDP[(j-1)*l1+(i-1)] + 1
				resultLen = max(resultLen, lcDP[j*l1+i])
				lastIndex = i
			} else {
				lcDP[j*l1+i] = 0
			}

		}
	}

	// convert to go runes and return the slice of those runes
	runes := []rune(s1)
	return string(runes[lastIndex : lastIndex+resultLen]), resultLen

}

func deduceCommonName(files []AnalyzedFile) (bool, string) {
	// TOOD impl
	return false, ""
}
