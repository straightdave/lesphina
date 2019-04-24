package lesphina

import (
	"fmt"
	"testing"
)

func TestParseAll(t *testing.T) {
	meta, err := parseSource("./test_fixture/test_source.go.t")
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	t.Logf("%v", meta.Json())

	if meta.NumFunction != len(meta.Functions) ||
		meta.NumInterface != len(meta.Interfaces) ||
		meta.NumStruct != len(meta.Structs) {
		t.Fail()
	}
}

func TestParseImports(t *testing.T) {
	meta, err := parseSource("./test_fixture/test_source.go.t")
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	for _, i := range meta.Imports {
		fmt.Printf("- %+v\n", i)
	}
}

func TestParsingInterfaces(t *testing.T) {
	meta, err := parseSource("./test_fixture/test_source.go.t")
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
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
	t.Logf("res: %+v", getArgs("t1"))
	t.Logf("res: %+v", getArgs("t1,t2"))
	t.Logf("res: %+v", getArgs("a t1, b t2"))

	// omit one type (not the last)
	t.Logf("res: %+v", getArgs("a1, a2 t1, b t2"))

	// omit any name will cause problems
	// but it's fine since that's illegal
	t.Logf("res: %+v", getArgs("t1, b t2"))
	t.Logf("res: %+v", getArgs("a t1, t2")) // same as: omit last type
}

func TestParsingEle(t *testing.T) {
	ele := &Element{
		Name:    "hahaha",
		RawType: "*XXXRequest",
	}

	parseEle(ele)

	t.Logf("parsed ele: %+v", ele)

	if ele.IsPointer != true {
		t.Fail()
	}
}

func TestJsonFieldName(t *testing.T) {
	ele := &Element{
		Name:   "MyEle",
		RawTag: "`" + `json:"haha,fasdfasdf"` + "`",
	}

	if ele.JsonFieldName() != "haha" {
		t.Fail()
	}
	t.Logf("%v", ele.JsonFieldName())

	ele2 := &Element{
		Name:   "MyEle",
		RawTag: "`" + `bson:"haha"` + "`",
	}

	if ele2.JsonFieldName() != "" {
		t.Fail()
	}
	t.Logf("%v", ele2.JsonFieldName())

}
