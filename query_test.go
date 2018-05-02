package lesphina

import (
	"testing"

	"github.com/straightdave/lesphina/entry"
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

	f := q.ByKind(entry.KindInterface).ByName("Int").ByName("Int0").First()
	if f == nil {
		t.Fail()
	}

	if f.GetName() != "Int0" {
		t.Fail()
	}

	ff, ok := f.(*entry.Interface) // '*entry.Interface' implements entry.Entry
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
