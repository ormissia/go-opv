// Copyright (c) 2021 安红豆. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package go_opv

import (
	"reflect"
	"strconv"
	"strings"
)

func isEmpty(value reflect.Value) bool {
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

func compareVerify(value reflect.Value, verifyStr, separator string) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice, reflect.Array:
		return compare(value.Len(), verifyStr, separator)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), verifyStr, separator)
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), verifyStr, separator)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), verifyStr, separator)
	default:
		return false
	}
}

func compare(value interface{}, verifyStr, separator string) bool {
	rule := strings.Split(verifyStr, separator)
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		length, err := strconv.ParseInt(rule[1], 10, 64)
		if err != nil {
			return false
		}
		switch {
		case rule[0] == eq:
			return val.Int() == length
		case rule[0] == ne:
			return val.Int() != length
		case rule[0] == gt:
			return val.Int() > length
		case rule[0] == lt:
			return val.Int() < length
		case rule[0] == ge:
			return val.Int() >= length
		case rule[0] == le:
			return val.Int() <= length
		default:
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		length, err := strconv.Atoi(rule[1])
		if err != nil {
			return false
		}
		switch {
		case rule[0] == eq:
			return val.Uint() == uint64(length)
		case rule[0] == ne:
			return val.Uint() != uint64(length)
		case rule[0] == gt:
			return val.Uint() > uint64(length)
		case rule[0] == lt:
			return val.Uint() < uint64(length)
		case rule[0] == ge:
			return val.Uint() >= uint64(length)
		case rule[0] == le:
			return val.Uint() <= uint64(length)
		default:
			return false
		}
	case reflect.Float32, reflect.Float64:
		length, err := strconv.ParseFloat(rule[1], 64)
		if err != nil {
			return false
		}
		switch {
		case rule[0] == eq:
			return val.Float() == length
		case rule[0] == ne:
			return val.Float() != length
		case rule[0] == gt:
			return val.Float() > length
		case rule[0] == lt:
			return val.Float() < length
		case rule[0] == ge:
			return val.Float() >= length
		case rule[0] == le:
			return val.Float() <= length
		default:
			return false
		}
	default:
		return false
	}
}
