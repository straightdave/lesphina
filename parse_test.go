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
