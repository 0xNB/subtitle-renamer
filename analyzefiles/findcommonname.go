// Package analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

// we use this function to deduce the longest common substring in a list of string
func lcSubStr(s1 string, s2 string) {
	// create two dim array that holds our longest common substrings for DP
	l1 := len(s1)
	l2 := len(s2)

	lcDP := make([]string, l1*l2)
	var resultLen int

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			
		}
	}

}

func deduceCommonName(files []AnalyzedFile) (bool, string) {

}
