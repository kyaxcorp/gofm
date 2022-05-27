package filemanager

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io/fs"
)

type NewFile struct {
	// File -> should be indicated, other params are optional!
	File fs.File
	// TODO : add here other methods of input like:
	//       - bytes
	//       - io

	//======= Optional ======= \\
	// Name -> can
	Name        string
	Description string
	CategoryID  *uuid.UUID
	// Extension -> is needed only when a physical file is not provided, but a Bytes input has been provided
	Extension string
	//======= Optional ======= \\

	// ========= Helpers =========\\
	// by indicating File, it will be read from there
	// TODO: should we add here the locations where to be saved...
	fileManager *FileManager
	// ========= Helpers =========\\

}

// NewFile -> create a new file in the DB
func (fm *FileManager) NewFile() *NewFile {
	return &NewFile{
		// Set the file manager as reference
		fileManager: fm,
	}
}

func (f *NewFile) db() *gorm.DB {
	return f.fileManager.db()
}
