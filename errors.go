package gofm

import "errors"

var (
	// =============== InputFile Manager Errors ===============\\

	ErrFmDatabaseClientError    = errors.New("file manager database client error") // TODO: should be renamed or adapted
	ErrFmFileNotFound           = errors.New("file manager file not found")
	ErrFmNewFileLocationMissing = errors.New("file manager file location missing")
	// =============== InputFile Manager Errors ===============\\

	//
)

// setError -> return the fmError for code simplicity
func (fm *FileManager) setError(fmError, intError error) error {
	fm.fmError = fmError
	fm.internalError = intError
	return fm.fmError
}
