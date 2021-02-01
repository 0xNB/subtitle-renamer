// Package renamefiles is responsible for renaming the previously found analyzed files.
package renamefiles

import (
	"fmt"
	"os"
)

// RenameFile renames a file and returns whether the operation succeeded or not
func RenameFile(oldName, newName string) bool {
	if error := os.Rename(oldName, newName); error != nil {
		fmt.Printf("Error renaming file %s->%s", oldName, newName)
		return false
	}
	return true
}
