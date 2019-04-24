package lesphina

import (
	"strings"
)

type Query struct {
	residue []Entry
}

// Query gets a query instance from lesphina instance.
func (les *Lesphina) Query() *Query {
	var set []Entry

	// flatten Meta
	// TODO: to support more kinds for now

	for _, v := range les.Meta.Vars {
		set = append(set, v)
	}

	for _, c := range les.Meta.Consts {
		set = append(set, c)
	}

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

	return &Query{residue: set}
}

// First resolves a query instance and returns the first qualified result.
func (q *Query) First() Entry {
	if len(q.residue) < 1 {
		return nil
	}

	return q.residue[0]
}

// All resolves a query instance and returns all qualified results.
func (q *Query) All() []Entry {
	return q.residue
}

// ByName filters entries by the type name.
// Supporting '~xxxx', 'xxxx~' and '~xxxx~' patterns
func (q *Query) ByName(name string) *Query {
	var res []Entry

	fPre := strings.HasPrefix(name, "~")
	fSuf := strings.HasSuffix(name, "~")
	name = strings.TrimLeft(name, "~")
	name = strings.TrimRight(name, "~")

	if fPre && fSuf {
		for _, e := range q.residue {
			if strings.Contains(e.GetName(), name) {
				res = append(res, e)
			}
		}
	} else if fPre {
		for _, e := range q.residue {
			if strings.HasSuffix(e.GetName(), name) {
				res = append(res, e)
			}
		}
	} else if fSuf {
		for _, e := range q.residue {
			if strings.HasPrefix(e.GetName(), name) {
				res = append(res, e)
			}
		}
	} else {
		for _, e := range q.residue {
			if e.GetName() == name {
				res = append(res, e)
			}
		}
	}

	q.residue = res
	return q
}

// ByKind filters entries by type kind.
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
