package filemanager

import (
	"github.com/google/uuid"
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


File Explorer
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

type File struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;not null;<-:create;default:gen_random_uuid()"`

	// File Name
	Name string `gorm:"size:255"`
	// FullName contains the Name+Extension
	FullName string `gorm:"size:255"`
	// It's the FullName of the file but containing the original name of it!
	OriginalName string `gorm:"size:255"`

	// File Description
	Description string `gorm:"size:255"`
	// Size Bytes
	Size int64
	// CategoryID -> also an optional field
	CategoryID uuid.UUID
	// Extension -> xls, doc etc...
	Extension string `gorm:"size:30"`
	// Type -> application/json, application/text, etc...
	Type string `gorm:"size:50"`

	// If it's related to some Module, it's just a reference
	// the user can save the id's of the files in other records or in a separate table...
	// yes, it's a performance degradation querying
	//RelatedToID *uuid.UUID
	//RelatedName string

	// these are the locations where the file is stored
	// there can be multiple ones, having as a backup option or for read performance...
	Locations []Location

	// TODO store meta data about the current file in the locations!
	// TODO: should we create a MetaLocation?! which will contain info about how is stored in the location

	// Here we store the reference to the file manager
	// TODO: should we make it private?!
	fileManager *FileManager

	//EncryptionPassword string
	//EncryptionAlgo     string

	CreatedAt *time.Time `gorm:"type:timestamptz;null;index:idx_core_dates"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;null;index:idx_core_dates"`
	DeletedAt *time.Time `gorm:"type:timestamptz;null;index:idx_core_dates"`

	CreatedByID *uuid.UUID `gorm:"type:uuid;null"`
	UpdatedByID *uuid.UUID `gorm:"type:uuid;null"`
	DeletedByID *uuid.UUID `gorm:"type:uuid;null"`
}

// TableName -> Get the Database table name from the file manager
func (f *File) TableName() string {
	return f.fileManager.GetFilesDBTableName()
}
