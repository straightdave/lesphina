package lesphina

import (
	"testing"
)

func TestParse(t *testing.T) {
	meta, err := parseSource("./test_fixture/test_source.go.t")
	if err != nil {
		t.Fatalf("parsing failed: %v\n", err)
	}

	t.Logf("\n%v\n", Jsonify(meta))

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
		t.Log(Jsonify(intr))
		t.Log("----")
	}
}
