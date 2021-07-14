package verify

import (
	"fmt"
	"reflect"
	"errors"
)

type Rules map[string][]string

const (
	NotEmpty = "NotEmpty"
)

func Verify(rqm interface{},roleMap Rules)(err error){
	typ := reflect.TypeOf(rqm)
	val := reflect.ValueOf(rqm)
	fmt.Println(typ)
	fmt.Println(val)
	kd := val.Kind()
	if kd != reflect.Struct {
		return errors.New("expect struct")
	}
	num := val.NumField()
	for i :=0;i<num;i++{
		itemType := typ.Field(i)
		itemValue := val.Field(i)
		if len(roleMap[itemType.Name]) >0{
			for _,v :=range roleMap[itemType.Name]{
				if v == "NotEmpty" && isBlank(itemValue){
					return errors.New(itemType.Name + "值不能为空")
				}
			}
		}
	}
	return nil
}


func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}