package lesphina

import (
	"testing"
)

func TestParse(t *testing.T) {
	meta, err := parseSource("./test_fixture/test_source.go.t")
	if err != nil {
		t.Fatalf("parsing failed: %v\n", err)
	}

	t.Logf("\n%v\n", meta.Json())

	if meta.NumFunction != uint(len(meta.Functions)) ||
		meta.NumInterface != uint(len(meta.Interfaces)) ||
		meta.NumStruct != uint(len(meta.Structs)) {
		t.Fail()
	}
}

func TestParsingInterfaces(t *testing.T) {
	meta, err := parseSource("./test_fixture/test_source.go.t")
	if err != nil {
		t.Fatalf("parsing failed: %v\n", err)
	}

	interfaces := meta.Interfaces
	if len(interfaces) < 1 {
		t.Fail()
	}

	for _, intr := range interfaces {
		t.Log(intr.Json())
		t.Log("----")
	}
}

func TestGetArgs(t *testing.T) {
	t.Logf("res: %+v\n", getArgs("t1"))
	t.Logf("res: %+v\n", getArgs("t1,t2"))
	t.Logf("res: %+v\n", getArgs("a t1, b t2"))

	// omit one type (not the last)
	t.Logf("res: %+v\n", getArgs("a1, a2 t1, b t2"))

	// omit any name will cause problems
	// but it's fine since that's illegal
	t.Logf("res: %+v\n", getArgs("t1, b t2"))
	t.Logf("res: %+v\n", getArgs("a t1, t2")) // same as: omit last type
}

func TestParsingEle(t *testing.T) {
	ele := &Element{
		Name:    "hahaha",
		RawType: "*XXXRequest",
	}

	parseEle(ele)

	t.Logf("parsed ele: %+v\n", ele)

	if ele.IsPointer != true {
		t.Fail()
	}
}
