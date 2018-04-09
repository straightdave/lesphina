package lesphina

import (
	"testing"
)

func TestQuerySyntax(t *testing.T) {
	q := &Query{}

	res := q.ByName("some").ByType(FUNCTION).First()
	t.Logf("%#+v", res)

	res2 := q.Which(func(c Cond) bool {
		return c.Name == "some"
	}).First()
	t.Logf("%#+v", res2)
}
