package lesphina

import (
// "reflect"
)

type Cond struct {
	Name string
	Type Type
}

type Query struct {
	meta *Meta
	cond []*Cond
}

func (les *Lesphina) Query() *Query {
	return &Query{meta: les.Meta}
}

func (q *Query) ByType(t Type) *Query {
	q.cond = append(q.cond, &Cond{Type: t})
	return q
}

func (q *Query) ByName(n string) *Query {
	q.cond = append(q.cond, &Cond{Name: n})
	return q
}

func (q *Query) resolve() {

}

type Result struct {
}

func (q *Query) First() *Result {

	return &Result{}
}

func (q *Query) All() []*Result {
	return []*Result{&Result{}}
}

func (q *Query) Which(func(Cond) bool) *Query {
	return q
}

type Type int

const (
	IMPORT Type = iota
	STRUCT
	INTERFACE
	FUNCTION
	VAR
	CONST
)
