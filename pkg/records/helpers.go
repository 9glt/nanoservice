package records

import (
	"reflect"
	"strings"
)

func getModelName(a interface{}) string {
	if val, ok := a.(reflect.Value); ok {
		return getModelName(val.Interface())
	}
	if val, ok := a.(*reflect.Value); ok {
		return getModelName(val.Elem().Interface())
	}
	if reflect.TypeOf(a).Kind() == reflect.Ptr {
		return getModelName(reflect.ValueOf(a).Elem().Interface())
	}
	if reflect.TypeOf(a).Kind() == reflect.Slice {
		return getModelName(reflect.New(reflect.TypeOf(a).Elem()))
	}
	//if val, ok := a.(reflect.Type); ok {
	//	return strings.ToLower(val.Name())
	//}
	return strings.ToLower(reflect.TypeOf(a).Name())
}

func CheckError(err error) error {
	// 		if err.Error() == "record not found" {
	if err != nil {
		if err.Error() != "record not found" {
			// log.Printf("DB error in Get(%s)-(%v). %s", getModelName(a), a, err.Error())
			return nil
		}
	}
	return err
}
