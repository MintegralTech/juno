package helpers

import (
	"fmt"
	"github.com/MintegralTech/juno/document"
	"strings"
)

const ACCURACY = 0.000001

type Comparable interface {
	Compare(a, b interface{}) int
}

type Func func(a, b interface{}) int

func (f Func) Compare(a, b interface{}) int {
	return f(a, b)
}

var intCompare Func = func(a, b interface{}) int {
	switch b.(type) {
	case int8:
		return int8Func(a.(int8), b.(int8))
	//case *int8:
	//	return int8Func(*(a.(*int8)), *(b.(*int8)))
	case int16:
		return int16Func(a.(int16), b.(int16))
	//case *int16:
	//	return int16Func(*(a.(*int16)), *(b.(*int16)))
	case int:
		return intFunc(a.(int), b.(int))
	//case *int:
	//	return intFunc(*(a.(*int)), *(b.(*int)))
	case int32:
		return int32Func(a.(int32), b.(int32))
	//case *int32:
	//	return int32Func(*(a.(*int32)), *(b.(*int32)))
	case int64:
		return int64Func(a.(int64), b.(int64))
	//case *int64:
	//	return int64Func(*(a.(*int64)), *(b.(*int64)))
	case byte:
		return byteFunc(a.(byte), b.(byte))
	//case *byte:
	//	return byteFunc(*(a.(*byte)), *(b.(*byte)))
	case uint16:
		return uint16Func(a.(uint16), b.(uint16))
	//case *uint16:
	//	return uint16Func(*(a.(*uint16)), *(b.(*uint16)))
	case uint32:
		return uint32Func(a.(uint32), b.(uint32))
	//case *uint32:
	//	return uint32Func(*(a.(*uint32)), *(b.(*uint32)))
	case uint:
		return uintFunc(a.(uint), b.(uint))
	//case *uint:
	//	return uintFunc(*(a.(*uint)), *(b.(*uint)))
	case uint64:
		return uint64Func(a.(uint64), b.(uint64))
	//case *uint64:
	//	return uint64Func(*(a.(*uint64)), *(b.(*uint64)))
	case document.DocId:
		return docIdFunc(a.(document.DocId), b.(document.DocId))
	default:
		panic(fmt.Sprintf("parameters[%v[%T] - %v[%T]] type wrong.", a, a, b, b))
	}
}

var floatCompare Func = func(a, b interface{}) int {
	switch b.(type) {
	case float32:
		return float32Func(a.(float32), b.(float32))
	//case *float32:
	//	return float32Func(*(a.(*float32)), *(b.(*float32)))
	case float64:
		return float64Func(a.(float64), b.(float64))
	//case *float64:
	//	return float64Func(*(a.(*float64)), *(b.(*float64)))
	default:
		panic(fmt.Sprintf("parameters[%v [%T] - %v [%T]] type wrong.", a, a, b, b))
	}
}

var stringCompare Func = func(a, b interface{}) int {
	switch b.(type) {
	case string:
		return stringFunc(a.(string), b.(string))
	//case *string:
	//	return stringFunc(*(a.(*string)), *(b.(*string)))
	default:
		panic(fmt.Sprintf("parameters[%v [%T] - %v [%T]] type wrong.", a, a, b, b))
	}
}

func int8Func(i, j int8) int {
	return int(i - j)
}

func int16Func(i, j int16) int {
	return int(i - j)
}

func int32Func(i, j int32) int {
	return int(i - j)
}

func int64Func(i, j int64) int {
	return int(i - j)
}

func intFunc(i, j int) int {
	return i - j
}

func byteFunc(i, j byte) int {
	return int(i - j)
}

func uint16Func(i, j uint16) int {
	return int(i - j)
}

func uint32Func(i, j uint32) int {
	return int(i - j)
}

func uintFunc(i, j uint) int {
	return int(i - j)
}

func uint64Func(i, j uint64) int {
	return int(i - j)
}

func float32Func(i, j float32) int {
	if j-i > ACCURACY {
		return -1
	} else if i-j > ACCURACY {
		return 1
	}
	return 0
}

func float64Func(i, j float64) int {
	if j-i > ACCURACY {
		return -1
	} else if i-j > ACCURACY {
		return 1
	}
	return 0
}

func stringFunc(i, j string) int {
	return strings.Compare(i, j)
}

func docIdFunc(i, j document.DocId) int {
	return int(i - j)
}
