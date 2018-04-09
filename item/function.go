package item

type Function struct {
	Name    string
	Recv    []*Element
	In      []*Element
	Out     []*Element
	RawBody string
}

type Element struct {
	Name    string
	RawType string
}
