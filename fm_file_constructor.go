package gofm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewFile struct {
	// InputFile -> should be indicated, other params are optional!
	InputFilePath string
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
	// by indicating InputFile, it will be read from there
	// TODO: should we add here the locations where to be saved...
	fileManager *FileManager
	// ========= Helpers =========\\

	//

	// ========= Locations ==========\\
	//Locations Locations
	Locations []NewFileLocation
}

type NewFileLocation struct {
	// this is the instance name
	LocationName string
	// add other options...

	// Path -> is optional
	Path string
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
