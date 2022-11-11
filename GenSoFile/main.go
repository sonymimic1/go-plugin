package main

import (
	"unsafe"
	"reflect"
	"github.com/sonymimic1/go-plugin/Iface/common"
)

type Gen struct{}

// used go mod build (-trimpath)
func Add_UsedModuel(x, y int) common.Total {
	sum := common.Total{
		A: x,
		B: y,
	}
	sum.Totalsum = sum.A + sum.B
	return sum
}


func Add_UsedVendor(x int, y int, extra interface{}) interface{} {
	sum := common.Total{
		A: x,
		B: y,
	}
	extraData :=(*common.InputExtraData)(unsafe.Pointer(reflect.ValueOf(extra).Pointer()))
	sum.Totalsum = sum.A + sum.B + extraData.A + extraData.B
	return &sum
}
