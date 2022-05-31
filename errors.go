package gofm

import "errors"

var (
	// =============== InputFile Manager Errors ===============\\

	ErrFmDatabaseClientError    = errors.New("file manager database client error")        // TODO: should be renamed or adapted
	ErrFmDBClientQueryFailed    = errors.New("file manager database client query failed") // TODO: should be renamed or adapted
	ErrFmFileNotFound           = errors.New("file manager file not found")
	ErrFmNewFileLocationMissing = errors.New("file manager file location missing")
	ErrFmNoInputFile            = errors.New("file manager no input file")
	// =============== InputFile Manager Errors ===============\\

	//
)

// setError -> return the fmError for code simplicity
func (fm *FileManager) setError(fmError, intError error) error {
	fm.fmError = fmError
	fm.internalError = intError
	return fm.fmError
}

func (fm *FileManager) GetInternalError() error {
	return fm.internalError
}

func (fm *FileManager) GetError() error {
	return fm.fmError
}
