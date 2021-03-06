package gofm

import (
	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"
)

type NewFile struct {
	// InputFile -> should be indicated, other params are optional!
	InputFilePath string
	GraphQLFile   *graphql.Upload

	// FileModel - asta e modelul din partea la client...
	FileModel interface{}

	// TODO : add here other methods of input like:
	//       - bytes
	//       - io

	//======= Optional ======= \\
	// Name -> can
	// TODO: parametrii acestia pot fi scosi.. pentru ca ei pot fi scrisi insasi in Model
	//Name        string
	//Description string
	//CategoryID  *uuid.UUID
	//// Extension -> is needed only when a physical file is not provided, but a Bytes input has been provided
	//Extension string
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
	// ========= Locations ==========\\

}

type NewFileLocation struct {
	// this is the instance name
	LocationName string
	// add other options...

	// if none of the lower params are indicated, it will take the current files basedir and it will
	// recreate on the destination location
	//========= Optional Params ===========\\
	// DirPath -> is optional, a new folder path is indicated, the file name remains the same!
	DirPath string

	// FilePath -> is optional, a new full file path is indicated (it includes as the dir path)
	FilePath string
	//========= Optional Params ===========\\

}

// NewFile -> create a new file in the DB
func (fm *FileManager) NewFile() *NewFile {
	return &NewFile{
		// Set the file manager as reference
		fileManager: fm,
	}
}

func (f *NewFile) db() *gorm.DB {
	return f.fileManager.DB()
}
