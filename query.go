package lesphina

import (
	"strings"
)

type Query struct {
	residue []Entry
}

func (les *Lesphina) Query() *Query {
	var set []Entry

	// flatten Meta
	for _, s := range les.Meta.Structs {
		set = append(set, s)
	}

	for _, f := range les.Meta.Functions {
		set = append(set, f)
	}

	for _, i := range les.Meta.Interfaces {
		set = append(set, i)
		for _, im := range i.Methods {
			set = append(set, im)
		}
	}

	return &Query{
		residue: set,
	}
}

func (q *Query) First() Entry {
	if len(q.residue) < 1 {
		return nil
	}

	return q.residue[0]
}

func (q *Query) All() []Entry {
	return q.residue
}

func (q *Query) ByName(name string) *Query {
	var res []Entry
	for _, e := range q.residue {
		if strings.Contains(e.GetName(), name) {
			res = append(res, e)
		}
	}
	q.residue = res
	return q
}

func (q *Query) ByKind(kind Kind) *Query {
	var res []Entry
	for _, e := range q.residue {
		if e.GetKind() == kind {
			res = append(res, e)
		}
	}
	q.residue = res
	return q
}
