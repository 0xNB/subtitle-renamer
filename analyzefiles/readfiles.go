// Package analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// FileType is a generic type for a file loaded from disk
type FileType int

const (
	// VID is a generic video type
	VID FileType= iota
	// SUB is a generic sub type
	SUB 
	// MKV files represent video formats
	MKV 
	// ASS file are a form of subtitle file
	ASS 
	// MP4 files represent video files but shouldn't be too common anymore
	MP4
	// SRT files represent subtitles.
	SRT
	// NIL type is an unrecognized file
	NIL
)

// AnalyzedFile can return its own type, the underlying file or name
type AnalyzedFile interface {
	FileName() string
	FileType() FileType
	FileInfo() os.FileInfo
}

// NilFile is a unrecognized file
type NilFile struct{}

// FileName for Nilfiles don't need a name
func (n NilFile) FileName() string { return "" }

// FileType for NilFiles is nil
func (n NilFile) FileType() FileType { return NIL }

// FileInfo for NilFiles is nil
func (n NilFile) FileInfo() os.FileInfo { return nil }

func containsString(strings []string, s string) bool {
	for _, cur := range strings {
		if cur == s {
			return true
		}
	}
	return false
}

// AnalyzedFileFromFileInfo generates a new AnalyzedFile from the given FileInformation is respect to supported video and sub formats
func AnalyzedFileFromFileInfo(info os.FileInfo) AnalyzedFile {

	fileExt := strings.ToLower(filepath.Ext(info.Name()))

	switch {
	case containsString(SupportedVideoFileTypes(), fileExt):
		// supported video file extension detected
		return VidFile{
			name:     info.Name(),
			fileInfo: info,
		}
	case containsString(SupportedSubFileTypes(), fileExt):
		// supported sub file extension detected
		return SubFile{
			Name:     info.Name(),
			OSFileInfo: info,
		}
	default:
		return NilFile{}
	}

}

// ReadFilesFromDir takes an absolute path or relative path and outputs a
func ReadFilesFromDir(folderPath string) []AnalyzedFile {

	scannedFiles := make([]AnalyzedFile, 0, 100)

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal("Couldn't read directory", err)
	}

	for _, file := range files {
		scannedFiles = append(scannedFiles, AnalyzedFileFromFileInfo(file))
	}

	return scannedFiles
}
