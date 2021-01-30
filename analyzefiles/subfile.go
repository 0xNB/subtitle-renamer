// The analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

import "os"

// Subfile identifies files that contain subtitles
type SubFile struct {
	name     string
	fileInfo os.FileInfo
}

// AnalyzedFileInterface
func (sub SubFile) FileType() string {
	return "sub"
}

func (sub SubFile) FileName() string {
	return sub.name
}

func (sub SubFile) FileInfo() os.FileInfo {
	return sub.fileInfo
}
