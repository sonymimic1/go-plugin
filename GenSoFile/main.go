package main

import (
	"github.com/sonymimic1/go-plugin/Iface/common"
)

// used go mod build (-trimpath)
func Add_UsedModuel(x, y int) common.Total {
	sum := common.Total{
		A: x,
		B: y,
	}
	sum.Totalsum = sum.A + sum.B
	return sum
}

// used go mod vendor build
func Add_UsedVendor(x, y int) interface{} {
	sum := common.Total{
		A: x,
		B: y,
	}
	sum.Totalsum = sum.A + sum.B
	return &sum
}
