package item

type Interface struct {
	Name    string             `json:"name"`
	RawBody string             `json:"raw_body"`
	Methods []*InterfaceMethod `json:"methods"`
}

type InterfaceMethod struct {
	Name    string     `json:"name"`
	RawType string     `json:"raw_type"`
	In      []*Element `json:"in,omitempty"`
	Out     []*Element `json:"out,omitempty"`
}
