package filemanager

import "errors"

var (
	// =============== File Manager Errors ===============\\

	ErrFmDatabaseClientError = errors.New("file manager database client error") // TODO: should be renamed or adapted
	ErrFmFileNotFound        = errors.New("file manager file not found")

	// =============== File Manager Errors ===============\\

	//
)

// setError -> return the fmError for code simplicity
func (fm *FileManager) setError(fmError, intError error) error {
	fm.fmError = fmError
	fm.internalError = intError
	return fm.fmError
}
