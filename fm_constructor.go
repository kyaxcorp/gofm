package filemanager

// NewFileManager -> creates a new instance of file manager which you can interact with
// Everytime you need to
/*func NewFileManager(fm ...*FileManager) *FileManager {
	// create a new instance which manages files

	var _fm *FileManager
	if len(fm) > 0 {
		_fm = fm[0]
	} else {
		_fm = &FileManager{}
	}

	return _fm
}*/

// GetInstance -> it should be called once somewhere in the app?!
// It checks for table structure
func GetInstance(fm ...*FileManager) *FileManager {
	// create a new instance which manages files

	var _fm *FileManager
	if len(fm) > 0 {
		_fm = fm[0]
	} else {
		_fm = &FileManager{}
	}

	if _fm.DBAutoMigrate {

	}

	return _fm
}

func (fm *FileManager) DBAutoMigrate() {
	// Migrate if necessary
	fm.DBClient.AutoMigrate(
		&File{},
		&Location{},
	)
}
