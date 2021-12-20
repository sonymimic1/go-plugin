package main

import (
	"fmt"
	"plugin"
	"reflect"
	"unsafe"

	"github.com/sonymimic1/go-plugin/Iface/common"
)

func main() {

	p, err := plugin.Open("GenSofile.so")
	if err != nil {
		panic(err)
	}

	// usd go mod vendor
	Add_UsedVendor, err2 := p.Lookup("Add_UsedVendor")
	if err2 != nil {
		panic(err)
	}
	sum := (*common.Total)(unsafe.Pointer(reflect.ValueOf(Add_UsedVendor.(func(int, int) interface{})(3, 1)).Pointer()))

	fmt.Printf("%v\n", sum.Totalsum)

}
