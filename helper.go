package gofm

import (
	"io/fs"
	"net/http"
	"strings"
)

type ReadFile interface {
	Read(b []byte) (int, error)
}

func GetFileContentType(f ReadFile) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := f.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

type FileMetaData struct {
	ContentType string
	Size        int64
	Name        string
	BaseName    string
	Extension   string
}

func GetFileMetaData(f fs.File) (*FileMetaData, error) {
	contentType, _err := GetFileContentType(f)
	if _err != nil {
		return nil, _err
	}
	fileInfo, _err := f.Stat()
	if _err != nil {
		return nil, _err
	}

	fileBaseName := fileInfo.Name()
	// Try getting the file name and extension

	extensionSep := "."
	var fileName, extension string
	if strings.Contains(fileBaseName, extensionSep) {
		fileBaseNameSplit := strings.Split(fileBaseName, extensionSep)
		// Recreate the file name without extension
		fileName = strings.Join(fileBaseNameSplit[:len(fileBaseNameSplit)-2], extensionSep)
		// take the last one!
		extension = fileBaseNameSplit[len(fileBaseNameSplit)-1]
	} else {
		// No extension available
		fileName = fileBaseName
	}

	return &FileMetaData{
		ContentType: contentType,
		Size:        fileInfo.Size(),
		Name:        fileName,
		BaseName:    fileBaseName,
		Extension:   extension,
	}, nil
}
