package filemanager

import "gorm.io/gorm"

// FileManager - metadata will be stored
//
type FileManager struct {
	// The DB Client
	DBClient      *gorm.DB
	DBAutoMigrate bool

	Name string
	// this is the database table name
	DBTableName string
	TablePrefix string
	// TODO: add prefix to the tables

	// TODO: check if this is ok to have!
	// fmError -> is the file manager error
	fmError error

	// TODO: check if this is ok to have!
	// internalError -> is raised by the used applications/drivers
	internalError error
}

// getDBTablePrefix -> return the database table prefix
func (fm *FileManager) getDBTablePrefix() string {
	if fm.TablePrefix != "" {
		return fm.TablePrefix + "_"
	}
	return ""
}

// GetFilesDBTableName - Here we store information about our files
func (fm *FileManager) GetFilesDBTableName() string {
	return fm.getDBTablePrefix() + "fm_files"
}

// GetLocationsDBTableName - here we store information a
func (fm *FileManager) GetLocationsDBTableName() string {
	return fm.getDBTablePrefix() + "fm_locations"
}
