package item

type Function struct {
	Name    string     `json:"name"`
	Recv    []*Element `json:"recv,omitempty"`
	In      []*Element `json:"in,omitempty"`
	Out     []*Element `json:"out,omitempty"`
	RawBody string     `json:"raw_body"`
}

func (i *Function) InParams() []*Element  { return i.In }
func (i *Function) OutParams() []*Element { return i.Out }

type Element struct {
	Name    string `json:"name"`
	RawType string `json:"raw_type"`

	IsPointer  bool `json:"is_pointer"`
	IsSlice    bool `json:"is_slice"`
	IsVariadic bool `json:"is_variadic"`
	IsMap      bool `json:"is_map"`

	BaseType  string `json:"base_type"`
	KeyType   string `json:"key_type,omitempty"`
	ValueType string `json:"value_type,omitempty"`
}
