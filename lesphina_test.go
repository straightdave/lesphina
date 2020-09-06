package lesphina

import (
	"testing"
)

func TestDumpRestore(t *testing.T) {
	les, err := Read("test_fixture/myapp.pb.gosrc")
	if err != nil || les == nil {
		t.Log(err.Error())
		t.FailNow()
	}

	// t.Log("meta:" + les.Meta.JSON())

	dump := les.DumpString()
	t.Log("dump:" + dump)

	les1 := Restore(dump)
	t.Log("restore:" + les1.Meta.JSON())
}

func TestDumpRestore2(t *testing.T) {
	les, err := Read("test_fixture/newversion.pb.gosrc")
	if err != nil || les == nil {
		t.Log(err.Error())
		t.FailNow()
	}

	t.Log("meta:" + les.Meta.JSON())

	dump := les.DumpString()
	t.Log("dump:" + dump)

	les1 := Restore(dump)
	t.Log("restore:" + les1.Meta.JSON())
}
