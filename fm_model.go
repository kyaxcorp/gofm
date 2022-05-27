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
