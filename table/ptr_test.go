package table

import (
	"fmt"
	"reflect"
	"testing"
)

type xint int

func TestSliceCast(t *testing.T) {
	//var x interface{} = []int{2, 3,}
	//y := reflection.Cast(x, reflect.TypeOf(make([]xint, 0)))
	//xints := y.([]xint)
	//fmt.Println(xints)
	xint := make([]xint, 0)
	xintType := reflect.TypeOf(xint)
	xintElem := xintType.Elem()
	fmt.Println(xintElem.Size())
	fmt.Println(xintElem.Kind()) // 转化为int
}

func TestValue(t *testing.T) {
	f := 1.234
	//x := reflect.ValueOf(f).Interface()
	fval := reflect.ValueOf(f)
	fmt.Println(fval.CanAddr()) // false
	fptr := reflect.ValueOf(&f)
	fval = fptr.Elem() // valueptr -> value
	fmt.Println(fval.CanAddr())
}

func TestReflect(t *testing.T) {
	ints := []int{1, 3, 5}
	intsV := reflect.ValueOf(ints)
	fmt.Println(intsV.CanAddr())
	fmt.Println(intsV.Index(2).CanAddr())
}

func TestZero(t *testing.T) {
	fmt.Println(0 / 1)
}
