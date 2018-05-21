package lesphina

import (
	"testing"
)

func TestDumpRestore(t *testing.T) {
	les, err := Read("test_fixture/myapp.pb.go")
	if err != nil || les == nil {
		t.Log(err.Error())
		t.FailNow()
	}

	// t.Log("meta:" + les.Meta.Json())

	dump := les.DumpString()
	t.Log("dump:" + dump)

	les1 := Restore(dump)
	t.Log("restore:" + les1.Meta.Json())
}
