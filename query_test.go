package lesphina

import (
	"testing"
)

func TestQuery(t *testing.T) {
	les, err := Read("./test_fixture/test_source.go.t")
	if err != nil {
		t.Fatalf("parsing failed: %v\n", err)
	}

	q := les.Query()
	if q == nil || len(q.residue) < 1 {
		t.Fail()
	}

	f := q.ByKind(KindInterface).ByName("~Int~").ByName("Int0").First()
	if f == nil {
		t.Fail()
	}

	if f.GetName() != "Int0" {
		t.Fail()
	}

	ff, ok := f.(*Interface) // '*Interface' implements Entry
	if !ok {
		t.Fail()
	}

	if len(ff.Methods) < 1 {
		t.Fail()
	}

	for _, m := range ff.Methods {
		t.Logf("method: %s\n", m.GetName())
	}
}

func TestByName(t *testing.T) {
	les, err := Read("./test_fixture/test_source.go.t")
	if err != nil {
		t.Fatalf("parsing failed: %v\n", err)
	}

	q := les.Query()
	if len(q.residue) < 1 {
		t.Fatalf("query no init entry")
	}

	fun := q.ByKind(KindFunction).ByName("Func~").First()
	if fun == nil {
		t.Logf("found no Func")
		t.FailNow()
	}

	ff := fun.(*Function)
	t.Log("func found:", ff.GetName())

	q = les.Query()
	f := q.ByName("FuncNotExist").All() // no panic
	if len(f) > 0 {
		t.Fail()
	}
}
