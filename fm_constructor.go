package filemanager

import "log"

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

	log.Println(_fm)

	if _fm.DBAutoMigrate {
		_fm.DatabaseAutoMigrate()
	}

	return _fm
}

// DatabaseAutoMigrate - create all necessary tables, alter,add columns
func (fm *FileManager) DatabaseAutoMigrate() {
	// Migrate if necessary
	fm.db().AutoMigrate(
		&File{},
		//&File{fileManager: fm},
		//&Location{fileManager: fm},
	)
}
