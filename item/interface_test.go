package item

import (
	"testing"
)

func TestFirstInParam(t *testing.T) {
	method := &InterfaceMethod{
		Name: "haha",
		In: []*Element{
			&Element{
				Name:     "inparam1",
				RawType:  "*XXXRequest",
				BaseType: "XXXRequest",
			},

			&Element{
				Name:     "inparam2",
				RawType:  "*YYYRequest",
				BaseType: "YYYRequest",
			},
		},
	}

	if FirstInParam(method, "~Request").Name != "inparam1" {
		t.Fail()
	}

	// try to invoke nil pointer panic
	t.Logf("hahah: %v\n", FirstInParam(method, "~NotExists").Name)
}

func TestWorksForFunctions(t *testing.T) {
	fun := &Function{
		Name: "func1",
		In: []*Element{
			&Element{
				Name:     "inparam1",
				RawType:  "*XXXRequest",
				BaseType: "XXXRequest",
			},
		},
	}

	if FirstInParam(fun, "~Request").Name != "inparam1" {
		t.Fail()
	}
}
