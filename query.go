package lesphina

type Query struct {
}

type Cond struct {
	Name string
	Type Type
}

type Result struct {
}

func (q *Query) ByType(t Type) *Query {
	return q
}

func (q *Query) ByName(n string) *Query {
	return q
}

func (q *Query) Which(func(Cond) bool) *Query {
	return q
}

func (q *Query) First() *Result {
	return &Result{}
}

func (q *Query) All() []*Result {
	return []*Result{&Result{}}
}
