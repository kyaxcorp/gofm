package gofm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	// InstanceName, partitioning can be setup!
	FMInstance string    `gorm:"primaryKey;type:varchar(50);not null;<-:create"`
	ID         uuid.UUID `gorm:"primaryKey;not null;<-:create;default:gen_random_uuid()"`

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
	Locations Locations

	// TODO store meta data about the current file in the locations!
	// TODO: should we create a MetaLocation?! which will contain info about how is stored in the location

	// Here we store the reference to the file manager
	fileManager *FileManager `gorm:"-"`

	//EncryptionPassword string
	//EncryptionAlgo     string

	CreatedAt *time.Time `gorm:"null;index:idx_core_dates;<-:create"`
	UpdatedAt *time.Time `gorm:"null;index:idx_core_dates"`
	DeletedAt *time.Time `gorm:"null;index:idx_core_dates"`

	CreatedByID *uuid.UUID `gorm:"null;<-:create"`
	UpdatedByID *uuid.UUID `gorm:"null"`
	DeletedByID *uuid.UUID `gorm:"null"`
}

// TableName -> Get the Database table name from the file manager
// https://gorm.io/docs/conventions.html -> NOTE TableName doesn’t allow dynamic name, its result will be cached for future, to use dynamic name, you can use Scopes, for example:
func (f *File) TableName() string {
	return GetFilesDBTableName()
	//return f.fileManager.GetFilesDBTableName()
}

func (f *File) db() *gorm.DB {
	return f.fileManager.db()
}
