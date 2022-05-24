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
	MkDir()
	DeleteDir()
}

const (
	SaveToDB   = 1
	SaveToDisk = 2
	// Add other Save locations later on...
)

type SaveLocation struct {
	Database bool
	Disk     bool
	// TODO: where can we add destination settings
}

type File struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;not null;<-:create;default:gen_random_uuid()"`

	// File Name
	Name string
	// File Description
	Description string
	// Size Bytes
	Size string
	// CategoryID -> also an optional field
	CategoryID uuid.UUID
	// Extension -> xls, doc etc...
	Extension string
	// application/json, application/text, etc...
	Type string
	// Physical path of the file on the disk
	Path string
	// If it's related to some Module, it's just a reference
	RelatedToID *uuid.UUID
	RelatedName string
	// Is saved physically or in database, if in database, then it will be stored in another table
	IsPhysical bool

	// these are the locations where the file is stored
	// there can be multiple ones, having as a backup option or for read performance...
	Locations []Location
	// TODO store meta data about the current file in the locations!
	// TODO: should we create a MetaLocation?! which will contain info about how is stored in the location

	// Here we store the reference to the file manager
	// TODO: should we make it private?!
	FileManager *FileManager

	//EncryptionPassword string
	//EncryptionAlgo     string

	CreatedAt *time.Time `gorm:"type:timestamptz;null;index:idx_core_dates"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;null;index:idx_core_dates"`

	IsDeleted bool       `gorm:"type:bool;not null;default:false;index:idx_core_bool"`
	DeletedAt *time.Time `gorm:"type:timestamptz;null;index:idx_core_dates"`

	CreatedByID *uuid.UUID `gorm:"type:uuid;null"`
	UpdatedByID *uuid.UUID `gorm:"type:uuid;null"`
	DeletedByID *uuid.UUID `gorm:"type:uuid;null"`
}

// TableName -> Get the Database table name from the file manager
func (f *File) TableName() string {
	return f.FileManager.GetFilesDBTableName()
}
