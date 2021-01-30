// The analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

import "os"

// File that represents an entry for this show
type VidFile struct {
	name     string
	fileInfo os.FileInfo
}

// AnalyzedFileInterface
func (vid VidFile) FileType() string {
	return "vid"
}

func (vid VidFile) FileName() string {
	return vid.name
}

func (vid VidFile) FileInfo() os.FileInfo {
	return vid.fileInfo
}
