package gofm

import (
	"github.com/google/uuid"
	gofileDriver "github.com/kyaxcorp/gofile/driver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

/*
Use Cases:
Create new file:
- Define a file name
- give as input a file with :
	- physical path:
	- input buffer
	- bytes chunks
- generate a new UUID
- register the file in the DB through GORM
- start copying the file to the destination (only 1 destination! support for multiple is not planned)
- there is a status completion for each destination
-
Move FIle:


InputFile Explorer
*/

/*
	Ideal ar fi ca file manager-ul acesta sa aiba mai multe functii si drivere...
	Drivers:
		- Disk
		- Database
		- Google Drive
		- Dropbox
		- Telegram
		- OneDrive
		- WebAPI
		- WebDav
		- SFTP
		- FTP
		-
		etc...
*/

type DriverFileInterface interface {
	Save()
	Copy()
	Move()
	Delete()
	Create()
	Touch()
	List()

	// Do we need directory manager?!
	// if it's called file manager, then it should be related only to files...
	// how they are saved, should not interest the user?!
	// well, sometimes the user will want to indicate where to save a specific file...
	// but that's other thing, it's not related to making a directory or deleting it...
	// the user should be interested in saving the files and getting them back! that's it!
	//MkDir()
	//DeleteDir()
}

func (UUID) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// use field.Tag, field.TagSettings gets field's tags
	// checkout https://github.com/go-gorm/gorm/blob/master/schema/field.go for all options

	// returns different database type based on driver name
	switch db.Dialector.Name() {
	case "mysql", "sqlite":
		return "varchar(16)"
	case "postgres":
		return "uuid"
	}
	return ""
}

type LocationName string

type File struct {
	// InstanceName, partitioning can be setup!
	FMInstance string `gorm:"primaryKey;size:50;not null;<-:create"`
	ID         UUID   `gorm:"primaryKey;not null;<-:create"`

	// File Name
	Name string `gorm:"size:256"`
	// FullName contains the Name+Extension
	FullName string `gorm:"size:256"`
	// It's the FullName of the file but containing the original name of it!
	OriginalName string `gorm:"size:256;<-:create"`

	// File Description
	Description string `gorm:"size:256"`
	// Size Bytes
	Size int64 `gorm:"<-:create"`
	// CategoryID -> also an optional field
	CategoryID *uuid.UUID
	// Extension -> xls, doc etc...
	Extension string `gorm:"size:30;<-:create"`
	// ContentType -> application/json, application/text, etc...
	ContentType string `gorm:"size:50;<-:create"`

	// If it's related to some Module, it's just a reference
	// the user can save the id's of the files in other records or in a separate table...
	// yes, it's a performance degradation querying
	//RelatedToID *uuid.UUID
	//RelatedName string

	// these are the locations where the file is stored
	// there can be multiple ones, having as a backup option or for read performance...
	//Locations Locations
	Locations FileLocationsMeta
	// We are using it on reading
	locationsIndexed map[LocationName]FileLocationMeta

	// files -> indexed by location name
	files map[LocationName]gofileDriver.FileInterface
	// defaultFile -> this is the first file found in the array
	defaultFile gofileDriver.FileInterface

	// TODO store meta data about the current file in the locations!
	// TODO: should we create a MetaLocation?! which will contain info about how is stored in the location

	// Here we store the reference to the file manager
	fileManager *FileManager `gorm:"-"`

	//EncryptionPassword string
	//EncryptionAlgo     string

	FileCreatedAt *time.Time `gorm:"null;index:idx_core_dates;<-:create"`
	FileUpdatedAt *time.Time `gorm:"null;index:idx_core_dates"`

	CreatedAt *time.Time `gorm:"null;index:idx_core_dates;<-:create"`
	UpdatedAt *time.Time `gorm:"null;index:idx_core_dates"`
	DeletedAt *time.Time `gorm:"null;index:idx_core_dates"`

	CreatedByID *uuid.UUID `gorm:"null;<-:create"`
	UpdatedByID *uuid.UUID `gorm:"null"`
	DeletedByID *uuid.UUID `gorm:"null"`
}

func (f *File) SetFileManager(fm *FileManager) {
	f.fileManager = fm
}

func (f *File) SetLocationsIndexed(locations map[LocationName]FileLocationMeta) {
	f.locationsIndexed = locations
}

func (f *File) SetFilesIndexed(files map[LocationName]gofileDriver.FileInterface) {
	f.files = files
}

func (f *File) SetDefaultFile(file gofileDriver.FileInterface) {
	f.defaultFile = file
}

// TableName -> Get the Database table name from the file manager
// https://gorm.io/docs/conventions.html -> NOTE TableName doesn’t allow dynamic name, its result will be cached for future, to use dynamic name, you can use Scopes, for example:
func (f *File) TableName() string {
	return GetFilesDBTableName()
	//return f.fileManager.GetFilesDBTableName()
}

func (f *File) db() *gorm.DB {
	return f.fileManager.DB()
}

// ReadFromLoc -> sa citeasca din anumita locatie...
func (f *File) GetFromLoc(locName string) (gofileDriver.FileInterface, bool) {
	if file, ok := f.files[LocationName(locName)]; ok {
		return file, ok
	}
	return nil, false
}

func (f *File) Get() gofileDriver.FileInterface {
	return f.defaultFile
}

func (f *File) Read() ([]byte, error) {
	if f.defaultFile == nil {
		return nil, ErrFmFileNotFound
	}
	return f.defaultFile.Read()
}
