package main

import (
	"fmt"
	"plugin"

	"github.com/sonymimic1/go-plugin/Iface/common"
)

func main() {

	p, err := plugin.Open("GenSofile.so")
	if err != nil {
		panic(err)
	}

	// usd go mod vendor
	Add_UsedModuel, err := p.Lookup("Add_UsedModuel")
	if err != nil {
		panic(err)
	}
	sum_UsedModuel := Add_UsedModuel.(func(int, int) common.Total)(1, 1)

	fmt.Printf("%v\n", sum_UsedModuel.Totalsum)

}
