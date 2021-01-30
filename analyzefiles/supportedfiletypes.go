// Package analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

func SupportedVideoFileTypes() []string {
	return []string{
		"mkv",
	}
}

func SupportedSubFileTypes() []string {
	return []string{
		"srt",
	}
}
