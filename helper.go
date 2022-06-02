package gofm

import (
	"errors"
	"github.com/google/uuid"
	"github.com/kyaxcorp/gofile/driver/filesystem/helper"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

func structFieldExists(obj reflect.Value, fieldName string) bool {
	fieldVal := obj.FieldByName(fieldName)
	if fieldVal == (reflect.Value{}) {
		return false
	}
	return true
}

func structGetFieldVal(obj reflect.Value, fieldName string) (interface{}, error) {
	fieldVal := obj.FieldByName(fieldName)
	if fieldVal == (reflect.Value{}) {
		return nil, errors.New("struct field doesn't exist")
	}
	return fieldVal.Interface(), nil
}

func structSetFieldVal(obj reflect.Value, fieldName string, fieldValue interface{}) error {
	if !structFieldExists(obj, fieldName) {
		return errors.New("struct field " + fieldName + " doesn't exist")
	}

	f := obj.FieldByName(fieldName)
	v := reflect.ValueOf(fieldValue)

	objFieldType := f.Type().String()
	valFieldType := v.Type().String()

	if f.CanSet() {

		if objFieldType == "uuid.UUID" && valFieldType == "*uuid.UUID" {
			realVal := fieldValue.(*uuid.UUID)
			v = reflect.ValueOf(*realVal)
			f.Set(v)
			return nil
		} else if objFieldType == "*uuid.UUID" && valFieldType == "uuid.UUID" {
			realVal := fieldValue.(uuid.UUID)
			v = reflect.ValueOf(&realVal)
			f.Set(v)
			return nil

		} else if objFieldType == "gofm.UUID" && valFieldType == "*gofm.UUID" {
			realVal := fieldValue.(*UUID)
			v = reflect.ValueOf(*realVal)
			f.Set(v)
			return nil
		} else if objFieldType == "*gofm.UUID" && valFieldType == "gofm.UUID" {
			realVal := fieldValue.(UUID)
			v = reflect.ValueOf(&realVal)
			f.Set(v)
			return nil
		} else if objFieldType == "time.Time" && valFieldType == "*time.Time" {
			realVal := fieldValue.(*time.Time)
			v = reflect.ValueOf(*realVal)
			f.Set(v)
			return nil
		} else if objFieldType == "*time.Time" && valFieldType == "time.Time" {
			realVal := fieldValue.(time.Time)
			v = reflect.ValueOf(&realVal)
			f.Set(v)
			return nil
		}

		f.Set(v)
		return nil
	}
	return errors.New("can't set field value for struct field " + fieldName)
}

type temporaryFileOptions struct {
	SelfDestruct   bool
	SelfDestructIn time.Duration
	FileName       string
	FileData       []byte
}

func generateTmpFile(o temporaryFileOptions) (string, error) {
	tmpFolderID, _err := uuid.NewRandom()
	if _err != nil {
		return "", _err
	}
	// create the temporary folder where to save the file
	tmpFolder := os.TempDir() + filepath.FromSlash("/") + tmpFolderID.String()
	if !helper.FolderExists(tmpFolder) {
		_err = helper.MkDir(tmpFolder, 0751)
		if _err != nil {
			return "", _err
		}
	}
	// Delete the folder and the files inside it after copying the destination locations
	if o.SelfDestruct {
		time.AfterFunc(o.SelfDestructIn, func() {
			helper.FolderDelete(tmpFolder)
		})
	}

	//defer helper.FolderDelete(tmpFolder)
	// Write the file to that generated folder
	tmpFileFullPath := tmpFolder + filepath.FromSlash("/") + o.FileName
	_err = os.WriteFile(tmpFileFullPath, o.FileData, 0751)
	if _err != nil {
		return "", _err
	}
	return tmpFileFullPath, nil
}
