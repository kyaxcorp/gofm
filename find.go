package filemanager

import "github.com/google/uuid"

type FindOptions struct {
	ID   uuid.UUID
	Name string
}

// FindFile -> it will search for the described file in the database
// if nothing is found, an error will be returned
func FindFile(o FindOptions) *File {

	return nil
}
