package entry

type Function struct {
	Name    string     `json:"name"`
	Recv    []*Element `json:"recv,omitempty"`
	In      []*Element `json:"in,omitempty"`
	Out     []*Element `json:"out,omitempty"`
	RawBody string     `json:"raw_body"`
}

func (i *Function) GetName() string                       { return i.Name }
func (i *Function) GetKind() Kind                         { return KindFunction }
func (i *Function) Json() string                          { return marshal(i) }
func (i *Function) InParams() []*Element                  { return i.In }
func (i *Function) OutParams() []*Element                 { return i.Out }
func (i *Function) FirstInParam(pattern string) *Element  { return firstInParam(i, pattern) }
func (i *Function) FirstOutParam(pattern string) *Element { return firstOutParam(i, pattern) }

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

func (i *Element) GetName() string { return i.Name }
func (i *Element) GetKind() Kind   { return KindElement }
func (i *Element) Json() string    { return marshal(i) }
