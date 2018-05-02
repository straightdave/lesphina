package entry

import (
	"testing"
)

func TestSomeToJson(t *testing.T) {
	obj := &Import{
		Alias: "ctx",
		Name:  "context",
	}

	if len(obj.Json()) < 3 {
		// in case "{}"
		t.Fail()
	}
}

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

	if len(method.Json()) < 3 {
		t.Fail()
	}

	if firstInParam(method, "~Request").Name != "inparam1" {
		t.Fail()
	}

	if method.FirstInParam("~Request").Name != "inparam1" {
		t.Fail()
	}

	// try to invoke nil pointer panic
	t.Logf("hahah: %v\n", firstInParam(method, "~NotExists").Name)
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

	if len(fun.Json()) < 3 {
		t.Fail()
	}

	if firstInParam(fun, "~Request").Name != "inparam1" {
		t.Fail()
	}

	if fun.FirstInParam("~Request").Name != "inparam1" {
		t.Fail()
	}
}
