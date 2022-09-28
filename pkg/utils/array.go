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

func RemoveDuplicateValues(intSlice []uint) []uint {
	keys := make(map[uint]bool)
	var list []uint
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
