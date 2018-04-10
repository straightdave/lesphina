package item

type Function struct {
	Name    string     `json:"name"`
	Recv    []*Element `json:"recv"`
	In      []*Element `json:"in"`
	Out     []*Element `json:"out"`
	RawBody string     `json:"raw_body"`
}

type Element struct {
	Name    string `json:"name"`
	RawType string `json:"raw_type"`
}
