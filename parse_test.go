package lesphina

import (
	"testing"
)

func TestParse(t *testing.T) {
	meta, err := parseSource("./test_fixture/test_source.go.t")
	if err != nil {
		t.Fatalf("parsing failed: %v\n", err)
	}

	t.Logf("%#+v\n", meta)

	// functions
	t.Logf("Func: %d in total\n", meta.NumFunction)
	for _, f := range meta.Functions {
		t.Logf("func# %s:\n", f.Name)
		t.Logf("recv=>\n")
		for _, r := range f.Recv {
			t.Logf("\t%s: %s\n", r.Name, r.RawType)
		}

		t.Logf("in-params=>\n")
		for _, in := range f.In {
			t.Logf("\t%s: %s\n", in.Name, in.RawType)
		}

		t.Logf("out-params=>\n")
		for _, o := range f.Out {
			t.Logf("\t%s: %s\n", o.Name, o.RawType)
		}

		t.Logf("raw body=>\n")
		t.Logf("\n%s\n", f.RawBody)
		t.Log()
	}
}
