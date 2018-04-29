package item

import (
	"testing"
)

func TestFirstInParamLike(t *testing.T) {
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

	if method.FirstInParamLike("~Request").Name != "inparam1" {
		t.Fail()
	}

	// try to invoke nil pointer panic
	t.Logf("hahah: %v\n", method.FirstInParamLike("~NotExists").Name)
}
