// Package analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

import "os"

// Subfile identifies files that contain subtitles
type SubFile struct {
	Name     string
	FullPath string
	OSFileInfo os.FileInfo
} 

// FileType AnalyzedFileInterface
func (sub SubFile) FileType() FileType {
	return SUB
}

func (sub SubFile) FileName() string {
	return sub.Name
}

func (sub SubFile) FileInfo() os.FileInfo {
	return sub.OSFileInfo
}
