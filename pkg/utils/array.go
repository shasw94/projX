package utils

import "reflect"

func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}
	return false
}

func JoinUintArrays(array ...[]uint) []uint {
	var j []uint
	for _, a := range array {
		for _, b := range a {
			j = append(j, b)
		}
	}
	return j
}

func JoinStringArrays(array ...[]string) []string {
	var j []string
	for _, a := range array {
		for _, b := range a {
			j = append(j, b)
		}
	}
	return j
}

func RemoveDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
