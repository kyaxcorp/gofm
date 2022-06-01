package gofm

import (
	"errors"
	"github.com/google/uuid"
	"reflect"
	"time"
)

func structFieldExists(obj reflect.Value, fieldName string) bool {
	fName := obj.FieldByName("Name")
	if fName == (reflect.Value{}) {
		return false
	}
	return true
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
