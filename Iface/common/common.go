package common

type (
	Total struct {
		A        int
		B        int
		Totalsum int
	}

	InputExtraData struct {
		A int
		B int
	}

	Icommon interface {
		Add_UsedVendor(int, int, interface{}) interface{}
	}
)
