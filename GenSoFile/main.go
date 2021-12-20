package main

import (
	"github.com/sonymimic1/go-plugin/Iface/common"
)

func Add(x, y int) interface{} {
	sum := common.Total{
		A: x,
		B: y,
	}
	sum.Totalsum = sum.A + sum.B
	return &sum
}
