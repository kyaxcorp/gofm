package gofm

import (
	"gorm.io/gorm"
)

var (
	// DBFilesTableName -> can be changed as needed...
	DBFilesTableName = "fm_files"
)

// FileManager - metadata will be stored
//
type FileManager struct {
	// The DB Client
	DBClient      *gorm.DB
	DBAutoMigrate bool

	// Name -> the name should be short max 50 chars
	Name string
	// this is the database table name
	//DBTableName string // TODO: what is it for?!
	//TablePrefix string
	// TODO: add prefix to the tables

	Locations        Locations
	LocationsIndexed map[string]Location

	// TODO: check if this is ok to have!
	// fmError -> is the file manager error
	fmError error

	// TODO: check if this is ok to have!
	// internalError -> is raised by the used applications/drivers
	internalError error
}

func GetFilesDBTableName() string {
	return DBFilesTableName
}

/*// getDBTablePrefix -> return the database table prefix
func (fm *FileManager) getDBTablePrefix() string {
	if fm.TablePrefix != "" {
		return fm.TablePrefix + "_"
	}
	return ""
}*/

/*// GetFilesDBTableName - Here we store information about our files
func (fm *FileManager) GetFilesDBTableName() string {
	log.Println("GetFilesDBTableName")
	return fm.getDBTablePrefix() + "fm_files"
}

// GetLocationsDBTableName - here we store information a
func (fm *FileManager) GetLocationsDBTableName() string {
	return fm.getDBTablePrefix() + "fm_locations"
}*/
